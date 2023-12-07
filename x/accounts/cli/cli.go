package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	v1 "cosmossdk.io/x/accounts/v1"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                name,
		Short:              "Transactions command for the " + name + " module",
		RunE:               client.ValidateCmd,
		DisableFlagParsing: true,
	}
	cmd.AddCommand(GetTxInitCmd(), GetExecuteCmd())
	return cmd
}

func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                name,
		Short:              "Query command for the " + name + " module",
		RunE:               client.ValidateCmd,
		DisableFlagParsing: true,
	}
	cmd.AddCommand(GetQueryAccountCmd())
	return cmd
}

func GetTxInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [account-type] [json-message]",
		Short: "Initialize a new account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()

			// we need to convert the message from json to a protobuf message
			// to know which message to use, we need to know the account type
			// init message schema.
			accClient := v1.NewQueryClient(clientCtx)
			schema, err := accClient.Schema(cmd.Context(), &v1.SchemaRequest{
				AccountType: args[0],
			})
			if err != nil {
				return err
			}

			msgBytes, err := encodeJSONToProto(schema.InitSchema.Request, args[1])
			if err != nil {
				return err
			}
			msg := v1.MsgInit{
				Sender:      sender.String(),
				AccountType: args[0],
				Message:     msgBytes,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetExecuteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute [account-address] [execute-msg-type-url] [json-message]",
		Short: "Execute state transition to account",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()

			schema, err := getSchemaForAccount(clientCtx, args[0])
			if err != nil {
				return err
			}

			msgBytes, err := handlerMsgBytes(schema.ExecuteHandlers, args[1], args[2])
			if err != nil {
				return err
			}
			msg := v1.MsgExecute{
				Sender:  sender.String(),
				Target:  args[0],
				Message: msgBytes,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetQueryAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [account-address] [query-request-type-url] [json-message]",
		Short: "Query account state",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			schema, err := getSchemaForAccount(clientCtx, args[0])
			if err != nil {
				return err
			}
			msgBytes, err := handlerMsgBytes(schema.QueryHandlers, args[1], args[2])
			if err != nil {
				return err
			}
			queryClient := v1.NewQueryClient(clientCtx)
			res, err := queryClient.AccountQuery(cmd.Context(), &v1.AccountQueryRequest{
				Target:  args[0],
				Request: msgBytes,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func getSchemaForAccount(clientCtx client.Context, addr string) (*v1.SchemaResponse, error) {
	queryClient := v1.NewQueryClient(clientCtx)
	accType, err := queryClient.AccountType(clientCtx.CmdContext, &v1.AccountTypeRequest{
		Address: addr,
	})
	if err != nil {
		return nil, err
	}
	return queryClient.Schema(clientCtx.CmdContext, &v1.SchemaRequest{
		AccountType: accType.AccountType,
	})
}

func handlerMsgBytes(handlersSchema []*v1.SchemaResponse_Handler, msgTypeURL, msgString string) (*codectypes.Any, error) {
	var msgSchema *v1.SchemaResponse_Handler
	for _, handler := range handlersSchema {
		if handler.Request == msgTypeURL {
			msgSchema = handler
			break
		}
	}
	if msgSchema == nil {
		return nil, fmt.Errorf("handler for message type %s not found", msgTypeURL)
	}
	return encodeJSONToProto(msgSchema.Request, msgString)
}

func encodeJSONToProto(name, jsonMsg string) (*codectypes.Any, error) {
	jsonBytes := []byte(jsonMsg)
	impl, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}
	msg := impl.New().Interface()
	err = protojson.Unmarshal(jsonBytes, msg)
	if err != nil {
		return nil, err
	}
	msgBytes, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return &codectypes.Any{
		TypeUrl: "/" + name,
		Value:   msgBytes,
	}, nil
}
