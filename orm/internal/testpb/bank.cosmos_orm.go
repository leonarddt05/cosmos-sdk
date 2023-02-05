// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package testpb

import (
	context "context"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type BalanceTable interface {
	Insert(ctx context.Context, balance *Balance) error
	Update(ctx context.Context, balance *Balance) error
	Save(ctx context.Context, balance *Balance) error
	Delete(ctx context.Context, balance *Balance) error
	Has(ctx context.Context, address string, denom string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, address string, denom string) (*Balance, error)
	List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error
	DeleteRange(ctx context.Context, from, to BalanceIndexKey) error

	doNotImplement()
}

type BalanceIterator struct {
	ormtable.Iterator
}

func (i BalanceIterator) Value() (*Balance, error) {
	var balance Balance
	err := i.UnmarshalMessage(&balance)
	return &balance, err
}

type BalanceIndexKey interface {
	id() uint32
	values() []interface{}
	balanceIndexKey()
}

// primary key starting index..
type BalancePrimaryKey = BalanceAddressDenomIndexKey

type BalanceAddressDenomIndexKey struct {
	vs []interface{}
}

func (x BalanceAddressDenomIndexKey) id() uint32            { return 0 }
func (x BalanceAddressDenomIndexKey) values() []interface{} { return x.vs }
func (x BalanceAddressDenomIndexKey) balanceIndexKey()      {}

func (this BalanceAddressDenomIndexKey) WithAddress(address string) BalanceAddressDenomIndexKey {
	this.vs = []interface{}{address}
	return this
}

func (this BalanceAddressDenomIndexKey) WithAddressDenom(address string, denom string) BalanceAddressDenomIndexKey {
	this.vs = []interface{}{address, denom}
	return this
}

type BalanceDenomIndexKey struct {
	vs []interface{}
}

func (x BalanceDenomIndexKey) id() uint32            { return 1 }
func (x BalanceDenomIndexKey) values() []interface{} { return x.vs }
func (x BalanceDenomIndexKey) balanceIndexKey()      {}

func (this BalanceDenomIndexKey) WithDenom(denom string) BalanceDenomIndexKey {
	this.vs = []interface{}{denom}
	return this
}

type balanceTable struct {
	table ormtable.Table
}

func (this balanceTable) Insert(ctx context.Context, balance *Balance) error {
	return this.table.Insert(ctx, balance)
}

func (this balanceTable) Update(ctx context.Context, balance *Balance) error {
	return this.table.Update(ctx, balance)
}

func (this balanceTable) Save(ctx context.Context, balance *Balance) error {
	return this.table.Save(ctx, balance)
}

func (this balanceTable) Delete(ctx context.Context, balance *Balance) error {
	return this.table.Delete(ctx, balance)
}

func (this balanceTable) Has(ctx context.Context, address string, denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, address, denom)
}

func (this balanceTable) Get(ctx context.Context, address string, denom string) (*Balance, error) {
	var balance Balance
	found, err := this.table.PrimaryKey().Get(ctx, &balance, address, denom)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &balance, nil
}

func (this balanceTable) List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceTable) ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceTable) DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this balanceTable) DeleteRange(ctx context.Context, from, to BalanceIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this balanceTable) doNotImplement() {}

var _ BalanceTable = balanceTable{}

func NewBalanceTable(db ormtable.Schema) (BalanceTable, error) {
	table := db.GetTable(&Balance{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Balance{}).ProtoReflect().Descriptor().FullName()))
	}
	return balanceTable{table}, nil
}

type SupplyTable interface {
	Insert(ctx context.Context, supply *Supply) error
	Update(ctx context.Context, supply *Supply) error
	Save(ctx context.Context, supply *Supply) error
	Delete(ctx context.Context, supply *Supply) error
	Has(ctx context.Context, denom string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, denom string) (*Supply, error)
	List(ctx context.Context, prefixKey SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error)
	ListRange(ctx context.Context, from, to SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error)
	DeleteBy(ctx context.Context, prefixKey SupplyIndexKey) error
	DeleteRange(ctx context.Context, from, to SupplyIndexKey) error

	doNotImplement()
}

type SupplyIterator struct {
	ormtable.Iterator
}

func (i SupplyIterator) Value() (*Supply, error) {
	var supply Supply
	err := i.UnmarshalMessage(&supply)
	return &supply, err
}

type SupplyIndexKey interface {
	id() uint32
	values() []interface{}
	supplyIndexKey()
}

// primary key starting index..
type SupplyPrimaryKey = SupplyDenomIndexKey

type SupplyDenomIndexKey struct {
	vs []interface{}
}

func (x SupplyDenomIndexKey) id() uint32            { return 0 }
func (x SupplyDenomIndexKey) values() []interface{} { return x.vs }
func (x SupplyDenomIndexKey) supplyIndexKey()       {}

func (this SupplyDenomIndexKey) WithDenom(denom string) SupplyDenomIndexKey {
	this.vs = []interface{}{denom}
	return this
}

type supplyTable struct {
	table ormtable.Table
}

func (this supplyTable) Insert(ctx context.Context, supply *Supply) error {
	return this.table.Insert(ctx, supply)
}

func (this supplyTable) Update(ctx context.Context, supply *Supply) error {
	return this.table.Update(ctx, supply)
}

func (this supplyTable) Save(ctx context.Context, supply *Supply) error {
	return this.table.Save(ctx, supply)
}

func (this supplyTable) Delete(ctx context.Context, supply *Supply) error {
	return this.table.Delete(ctx, supply)
}

func (this supplyTable) Has(ctx context.Context, denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, denom)
}

func (this supplyTable) Get(ctx context.Context, denom string) (*Supply, error) {
	var supply Supply
	found, err := this.table.PrimaryKey().Get(ctx, &supply, denom)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &supply, nil
}

func (this supplyTable) List(ctx context.Context, prefixKey SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return SupplyIterator{it}, err
}

func (this supplyTable) ListRange(ctx context.Context, from, to SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return SupplyIterator{it}, err
}

func (this supplyTable) DeleteBy(ctx context.Context, prefixKey SupplyIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this supplyTable) DeleteRange(ctx context.Context, from, to SupplyIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this supplyTable) doNotImplement() {}

var _ SupplyTable = supplyTable{}

func NewSupplyTable(db ormtable.Schema) (SupplyTable, error) {
	table := db.GetTable(&Supply{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Supply{}).ProtoReflect().Descriptor().FullName()))
	}
	return supplyTable{table}, nil
}

type BankStore interface {
	BalanceTable() BalanceTable
	SupplyTable() SupplyTable

	BankQueryServiceServer

	doNotImplement()
}

type bankStore struct {
	balance BalanceTable
	supply  SupplyTable
}

func (x bankStore) BalanceTable() BalanceTable {
	return x.balance
}

func (x bankStore) SupplyTable() SupplyTable {
	return x.supply
}

func (bankStore) doNotImplement() {}

var _ BankStore = bankStore{}

func NewBankStore(db ormtable.Schema) (BankStore, error) {
	balanceTable, err := NewBalanceTable(db)
	if err != nil {
		return nil, err
	}

	supplyTable, err := NewSupplyTable(db)
	if err != nil {
		return nil, err
	}

	return bankStore{
		balanceTable,
		supplyTable,
	}, nil
}
