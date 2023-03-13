// Code generated by pggen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

// Querier is a typesafe Go interface backed by SQL queries.
//
// Methods ending with Batch enqueue a query to run later in a pgx.Batch. After
// calling SendBatch on pgx.Conn, pgxpool.Pool, or pgx.Tx, use the Scan methods
// to parse the results.
type Querier interface {
	GetProfile(ctx context.Context, accountID int) (GetProfileRow, error)
	// GetProfileBatch enqueues a GetProfile query into batch to be executed
	// later by the batch.
	GetProfileBatch(batch genericBatch, accountID int)
	// GetProfileScan scans the result of an executed GetProfileBatch query.
	GetProfileScan(results pgx.BatchResults) (GetProfileRow, error)

	CreateProfile(ctx context.Context, accountID int, name string) (pgconn.CommandTag, error)
	// CreateProfileBatch enqueues a CreateProfile query into batch to be executed
	// later by the batch.
	CreateProfileBatch(batch genericBatch, accountID int, name string)
	// CreateProfileScan scans the result of an executed CreateProfileBatch query.
	CreateProfileScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	UpdateAlias(ctx context.Context, alias string, accountID int) (pgconn.CommandTag, error)
	// UpdateAliasBatch enqueues a UpdateAlias query into batch to be executed
	// later by the batch.
	UpdateAliasBatch(batch genericBatch, alias string, accountID int)
	// UpdateAliasScan scans the result of an executed UpdateAliasBatch query.
	UpdateAliasScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	UpdateName(ctx context.Context, name string, accountID int) (pgconn.CommandTag, error)
	// UpdateNameBatch enqueues a UpdateName query into batch to be executed
	// later by the batch.
	UpdateNameBatch(batch genericBatch, name string, accountID int)
	// UpdateNameScan scans the result of an executed UpdateNameBatch query.
	UpdateNameScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	UpdateProfilePicture(ctx context.Context, pictureUrl string, accountID int) (pgconn.CommandTag, error)
	// UpdateProfilePictureBatch enqueues a UpdateProfilePicture query into batch to be executed
	// later by the batch.
	UpdateProfilePictureBatch(batch genericBatch, pictureUrl string, accountID int)
	// UpdateProfilePictureScan scans the result of an executed UpdateProfilePictureBatch query.
	UpdateProfilePictureScan(results pgx.BatchResults) (pgconn.CommandTag, error)
}

type DBQuerier struct {
	conn  genericConn   // underlying Postgres transport to use
	types *typeResolver // resolve types by name
}

var _ Querier = &DBQuerier{}

// genericConn is a connection to a Postgres database. This is usually backed by
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
type genericConn interface {
	// Query executes sql with args. If there is an error the returned Rows will
	// be returned in an error state. So it is allowed to ignore the error
	// returned from Query and handle it in Rows.
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// QueryRow is a convenience wrapper over Query. Any error that occurs while
	// querying is deferred until calling Scan on the returned Row. That Row will
	// error with pgx.ErrNoRows if no rows are returned.
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	// Exec executes sql. sql can be either a prepared statement name or an SQL
	// string. arguments should be referenced positionally from the sql string
	// as $1, $2, etc.
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// genericBatch batches queries to send in a single network request to a
// Postgres server. This is usually backed by *pgx.Batch.
type genericBatch interface {
	// Queue queues a query to batch b. query can be an SQL query or the name of a
	// prepared statement. See Queue on *pgx.Batch.
	Queue(query string, arguments ...interface{})
}

// NewQuerier creates a DBQuerier that implements Querier. conn is typically
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
func NewQuerier(conn genericConn) *DBQuerier {
	return NewQuerierConfig(conn, QuerierConfig{})
}

type QuerierConfig struct {
	// DataTypes contains pgtype.Value to use for encoding and decoding instead
	// of pggen-generated pgtype.ValueTranscoder.
	//
	// If OIDs are available for an input parameter type and all of its
	// transitive dependencies, pggen will use the binary encoding format for
	// the input parameter.
	DataTypes []pgtype.DataType
}

// NewQuerierConfig creates a DBQuerier that implements Querier with the given
// config. conn is typically *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
func NewQuerierConfig(conn genericConn, cfg QuerierConfig) *DBQuerier {
	return &DBQuerier{conn: conn, types: newTypeResolver(cfg.DataTypes)}
}

// WithTx creates a new DBQuerier that uses the transaction to run all queries.
func (q *DBQuerier) WithTx(tx pgx.Tx) (*DBQuerier, error) {
	return &DBQuerier{conn: tx}, nil
}

// preparer is any Postgres connection transport that provides a way to prepare
// a statement, most commonly *pgx.Conn.
type preparer interface {
	Prepare(ctx context.Context, name, sql string) (sd *pgconn.StatementDescription, err error)
}

// PrepareAllQueries executes a PREPARE statement for all pggen generated SQL
// queries in querier files. Typical usage is as the AfterConnect callback
// for pgxpool.Config
//
// pgx will use the prepared statement if available. Calling PrepareAllQueries
// is an optional optimization to avoid a network round-trip the first time pgx
// runs a query if pgx statement caching is enabled.
func PrepareAllQueries(ctx context.Context, p preparer) error {
	if _, err := p.Prepare(ctx, getProfileSQL, getProfileSQL); err != nil {
		return fmt.Errorf("prepare query 'GetProfile': %w", err)
	}
	if _, err := p.Prepare(ctx, createProfileSQL, createProfileSQL); err != nil {
		return fmt.Errorf("prepare query 'CreateProfile': %w", err)
	}
	if _, err := p.Prepare(ctx, updateAliasSQL, updateAliasSQL); err != nil {
		return fmt.Errorf("prepare query 'UpdateAlias': %w", err)
	}
	if _, err := p.Prepare(ctx, updateNameSQL, updateNameSQL); err != nil {
		return fmt.Errorf("prepare query 'UpdateName': %w", err)
	}
	if _, err := p.Prepare(ctx, updateProfilePictureSQL, updateProfilePictureSQL); err != nil {
		return fmt.Errorf("prepare query 'UpdateProfilePicture': %w", err)
	}
	return nil
}

// typeResolver looks up the pgtype.ValueTranscoder by Postgres type name.
type typeResolver struct {
	connInfo *pgtype.ConnInfo // types by Postgres type name
}

func newTypeResolver(types []pgtype.DataType) *typeResolver {
	ci := pgtype.NewConnInfo()
	for _, typ := range types {
		if txt, ok := typ.Value.(textPreferrer); ok && typ.OID != unknownOID {
			typ.Value = txt.ValueTranscoder
		}
		ci.RegisterDataType(typ)
	}
	return &typeResolver{connInfo: ci}
}

// findValue find the OID, and pgtype.ValueTranscoder for a Postgres type name.
func (tr *typeResolver) findValue(name string) (uint32, pgtype.ValueTranscoder, bool) {
	typ, ok := tr.connInfo.DataTypeForName(name)
	if !ok {
		return 0, nil, false
	}
	v := pgtype.NewValue(typ.Value)
	return typ.OID, v.(pgtype.ValueTranscoder), true
}

// setValue sets the value of a ValueTranscoder to a value that should always
// work and panics if it fails.
func (tr *typeResolver) setValue(vt pgtype.ValueTranscoder, val interface{}) pgtype.ValueTranscoder {
	if err := vt.Set(val); err != nil {
		panic(fmt.Sprintf("set ValueTranscoder %T to %+v: %s", vt, val, err))
	}
	return vt
}

const getProfileSQL = `SELECT *
FROM profile
WHERE account_id = $1;`

type GetProfileRow struct {
	AccountID  *int             `json:"account_id"`
	Name       string           `json:"name"`
	Alias      *string          `json:"alias"`
	PictureUrl *string          `json:"picture_url"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

// GetProfile implements Querier.GetProfile.
func (q *DBQuerier) GetProfile(ctx context.Context, accountID int) (GetProfileRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetProfile")
	row := q.conn.QueryRow(ctx, getProfileSQL, accountID)
	var item GetProfileRow
	if err := row.Scan(&item.AccountID, &item.Name, &item.Alias, &item.PictureUrl, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return item, fmt.Errorf("query GetProfile: %w", err)
	}
	return item, nil
}

// GetProfileBatch implements Querier.GetProfileBatch.
func (q *DBQuerier) GetProfileBatch(batch genericBatch, accountID int) {
	batch.Queue(getProfileSQL, accountID)
}

// GetProfileScan implements Querier.GetProfileScan.
func (q *DBQuerier) GetProfileScan(results pgx.BatchResults) (GetProfileRow, error) {
	row := results.QueryRow()
	var item GetProfileRow
	if err := row.Scan(&item.AccountID, &item.Name, &item.Alias, &item.PictureUrl, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return item, fmt.Errorf("scan GetProfileBatch row: %w", err)
	}
	return item, nil
}

const createProfileSQL = `INSERT INTO profile(account_id, name)
VALUES (
        $1,
        $2
        );`

// CreateProfile implements Querier.CreateProfile.
func (q *DBQuerier) CreateProfile(ctx context.Context, accountID int, name string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateProfile")
	cmdTag, err := q.conn.Exec(ctx, createProfileSQL, accountID, name)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query CreateProfile: %w", err)
	}
	return cmdTag, err
}

// CreateProfileBatch implements Querier.CreateProfileBatch.
func (q *DBQuerier) CreateProfileBatch(batch genericBatch, accountID int, name string) {
	batch.Queue(createProfileSQL, accountID, name)
}

// CreateProfileScan implements Querier.CreateProfileScan.
func (q *DBQuerier) CreateProfileScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec CreateProfileBatch: %w", err)
	}
	return cmdTag, err
}

const updateAliasSQL = `UPDATE profile
SET alias = $1
WHERE account_id = $2;`

// UpdateAlias implements Querier.UpdateAlias.
func (q *DBQuerier) UpdateAlias(ctx context.Context, alias string, accountID int) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateAlias")
	cmdTag, err := q.conn.Exec(ctx, updateAliasSQL, alias, accountID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateAlias: %w", err)
	}
	return cmdTag, err
}

// UpdateAliasBatch implements Querier.UpdateAliasBatch.
func (q *DBQuerier) UpdateAliasBatch(batch genericBatch, alias string, accountID int) {
	batch.Queue(updateAliasSQL, alias, accountID)
}

// UpdateAliasScan implements Querier.UpdateAliasScan.
func (q *DBQuerier) UpdateAliasScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateAliasBatch: %w", err)
	}
	return cmdTag, err
}

const updateNameSQL = `UPDATE profile
SET name = $1
WHERE account_id = $2;`

// UpdateName implements Querier.UpdateName.
func (q *DBQuerier) UpdateName(ctx context.Context, name string, accountID int) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateName")
	cmdTag, err := q.conn.Exec(ctx, updateNameSQL, name, accountID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateName: %w", err)
	}
	return cmdTag, err
}

// UpdateNameBatch implements Querier.UpdateNameBatch.
func (q *DBQuerier) UpdateNameBatch(batch genericBatch, name string, accountID int) {
	batch.Queue(updateNameSQL, name, accountID)
}

// UpdateNameScan implements Querier.UpdateNameScan.
func (q *DBQuerier) UpdateNameScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateNameBatch: %w", err)
	}
	return cmdTag, err
}

const updateProfilePictureSQL = `UPDATE profile
SET picture_url = $1
WHERE account_id = $2;`

// UpdateProfilePicture implements Querier.UpdateProfilePicture.
func (q *DBQuerier) UpdateProfilePicture(ctx context.Context, pictureUrl string, accountID int) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateProfilePicture")
	cmdTag, err := q.conn.Exec(ctx, updateProfilePictureSQL, pictureUrl, accountID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateProfilePicture: %w", err)
	}
	return cmdTag, err
}

// UpdateProfilePictureBatch implements Querier.UpdateProfilePictureBatch.
func (q *DBQuerier) UpdateProfilePictureBatch(batch genericBatch, pictureUrl string, accountID int) {
	batch.Queue(updateProfilePictureSQL, pictureUrl, accountID)
}

// UpdateProfilePictureScan implements Querier.UpdateProfilePictureScan.
func (q *DBQuerier) UpdateProfilePictureScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateProfilePictureBatch: %w", err)
	}
	return cmdTag, err
}

// textPreferrer wraps a pgtype.ValueTranscoder and sets the preferred encoding
// format to text instead binary (the default). pggen uses the text format
// when the OID is unknownOID because the binary format requires the OID.
// Typically occurs if the results from QueryAllDataTypes aren't passed to
// NewQuerierConfig.
type textPreferrer struct {
	pgtype.ValueTranscoder
	typeName string
}

// PreferredParamFormat implements pgtype.ParamFormatPreferrer.
func (t textPreferrer) PreferredParamFormat() int16 { return pgtype.TextFormatCode }

func (t textPreferrer) NewTypeValue() pgtype.Value {
	return textPreferrer{ValueTranscoder: pgtype.NewValue(t.ValueTranscoder).(pgtype.ValueTranscoder), typeName: t.typeName}
}

func (t textPreferrer) TypeName() string {
	return t.typeName
}

// unknownOID means we don't know the OID for a type. This is okay for decoding
// because pgx call DecodeText or DecodeBinary without requiring the OID. For
// encoding parameters, pggen uses textPreferrer if the OID is unknown.
const unknownOID = 0