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
	GetResourceByID(ctx context.Context, resourceID int) (GetResourceByIDRow, error)
	// GetResourceByIDBatch enqueues a GetResourceByID query into batch to be executed
	// later by the batch.
	GetResourceByIDBatch(batch genericBatch, resourceID int)
	// GetResourceByIDScan scans the result of an executed GetResourceByIDBatch query.
	GetResourceByIDScan(results pgx.BatchResults) (GetResourceByIDRow, error)

	GetResources(ctx context.Context) ([]GetResourcesRow, error)
	// GetResourcesBatch enqueues a GetResources query into batch to be executed
	// later by the batch.
	GetResourcesBatch(batch genericBatch)
	// GetResourcesScan scans the result of an executed GetResourcesBatch query.
	GetResourcesScan(results pgx.BatchResults) ([]GetResourcesRow, error)

	DeleteResourceByID(ctx context.Context, resourceID int) (pgconn.CommandTag, error)
	// DeleteResourceByIDBatch enqueues a DeleteResourceByID query into batch to be executed
	// later by the batch.
	DeleteResourceByIDBatch(batch genericBatch, resourceID int)
	// DeleteResourceByIDScan scans the result of an executed DeleteResourceByIDBatch query.
	DeleteResourceByIDScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	InsertResource(ctx context.Context, params InsertResourceParams) (pgconn.CommandTag, error)
	// InsertResourceBatch enqueues a InsertResource query into batch to be executed
	// later by the batch.
	InsertResourceBatch(batch genericBatch, params InsertResourceParams)
	// InsertResourceScan scans the result of an executed InsertResourceBatch query.
	InsertResourceScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	UpdateResource(ctx context.Context, params UpdateResourceParams) (pgconn.CommandTag, error)
	// UpdateResourceBatch enqueues a UpdateResource query into batch to be executed
	// later by the batch.
	UpdateResourceBatch(batch genericBatch, params UpdateResourceParams)
	// UpdateResourceScan scans the result of an executed UpdateResourceBatch query.
	UpdateResourceScan(results pgx.BatchResults) (pgconn.CommandTag, error)
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
	if _, err := p.Prepare(ctx, getResourceByIDSQL, getResourceByIDSQL); err != nil {
		return fmt.Errorf("prepare query 'GetResourceByID': %w", err)
	}
	if _, err := p.Prepare(ctx, getResourcesSQL, getResourcesSQL); err != nil {
		return fmt.Errorf("prepare query 'GetResources': %w", err)
	}
	if _, err := p.Prepare(ctx, deleteResourceByIDSQL, deleteResourceByIDSQL); err != nil {
		return fmt.Errorf("prepare query 'DeleteResourceByID': %w", err)
	}
	if _, err := p.Prepare(ctx, insertResourceSQL, insertResourceSQL); err != nil {
		return fmt.Errorf("prepare query 'InsertResource': %w", err)
	}
	if _, err := p.Prepare(ctx, updateResourceSQL, updateResourceSQL); err != nil {
		return fmt.Errorf("prepare query 'UpdateResource': %w", err)
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

const getResourceByIDSQL = `SELECT *
FROM resource
WHERE resource_id = $1;`

type GetResourceByIDRow struct {
	ResourceID  int              `json:"resource_id"`
	Title       *string          `json:"title"`
	Description *string          `json:"description"`
	Link        *string          `json:"link"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	CategoryID  *int             `json:"category_id"`
}

// GetResourceByID implements Querier.GetResourceByID.
func (q *DBQuerier) GetResourceByID(ctx context.Context, resourceID int) (GetResourceByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetResourceByID")
	row := q.conn.QueryRow(ctx, getResourceByIDSQL, resourceID)
	var item GetResourceByIDRow
	if err := row.Scan(&item.ResourceID, &item.Title, &item.Description, &item.Link, &item.CreatedAt, &item.UpdatedAt, &item.CategoryID); err != nil {
		return item, fmt.Errorf("query GetResourceByID: %w", err)
	}
	return item, nil
}

// GetResourceByIDBatch implements Querier.GetResourceByIDBatch.
func (q *DBQuerier) GetResourceByIDBatch(batch genericBatch, resourceID int) {
	batch.Queue(getResourceByIDSQL, resourceID)
}

// GetResourceByIDScan implements Querier.GetResourceByIDScan.
func (q *DBQuerier) GetResourceByIDScan(results pgx.BatchResults) (GetResourceByIDRow, error) {
	row := results.QueryRow()
	var item GetResourceByIDRow
	if err := row.Scan(&item.ResourceID, &item.Title, &item.Description, &item.Link, &item.CreatedAt, &item.UpdatedAt, &item.CategoryID); err != nil {
		return item, fmt.Errorf("scan GetResourceByIDBatch row: %w", err)
	}
	return item, nil
}

const getResourcesSQL = `SELECT *
FROM resource;`

type GetResourcesRow struct {
	ResourceID  *int             `json:"resource_id"`
	Title       *string          `json:"title"`
	Description *string          `json:"description"`
	Link        *string          `json:"link"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	CategoryID  *int             `json:"category_id"`
}

// GetResources implements Querier.GetResources.
func (q *DBQuerier) GetResources(ctx context.Context) ([]GetResourcesRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetResources")
	rows, err := q.conn.Query(ctx, getResourcesSQL)
	if err != nil {
		return nil, fmt.Errorf("query GetResources: %w", err)
	}
	defer rows.Close()
	items := []GetResourcesRow{}
	for rows.Next() {
		var item GetResourcesRow
		if err := rows.Scan(&item.ResourceID, &item.Title, &item.Description, &item.Link, &item.CreatedAt, &item.UpdatedAt, &item.CategoryID); err != nil {
			return nil, fmt.Errorf("scan GetResources row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetResources rows: %w", err)
	}
	return items, err
}

// GetResourcesBatch implements Querier.GetResourcesBatch.
func (q *DBQuerier) GetResourcesBatch(batch genericBatch) {
	batch.Queue(getResourcesSQL)
}

// GetResourcesScan implements Querier.GetResourcesScan.
func (q *DBQuerier) GetResourcesScan(results pgx.BatchResults) ([]GetResourcesRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query GetResourcesBatch: %w", err)
	}
	defer rows.Close()
	items := []GetResourcesRow{}
	for rows.Next() {
		var item GetResourcesRow
		if err := rows.Scan(&item.ResourceID, &item.Title, &item.Description, &item.Link, &item.CreatedAt, &item.UpdatedAt, &item.CategoryID); err != nil {
			return nil, fmt.Errorf("scan GetResourcesBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetResourcesBatch rows: %w", err)
	}
	return items, err
}

const deleteResourceByIDSQL = `DELETE
FROM resource
WHERE resource_id = $1;`

// DeleteResourceByID implements Querier.DeleteResourceByID.
func (q *DBQuerier) DeleteResourceByID(ctx context.Context, resourceID int) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteResourceByID")
	cmdTag, err := q.conn.Exec(ctx, deleteResourceByIDSQL, resourceID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteResourceByID: %w", err)
	}
	return cmdTag, err
}

// DeleteResourceByIDBatch implements Querier.DeleteResourceByIDBatch.
func (q *DBQuerier) DeleteResourceByIDBatch(batch genericBatch, resourceID int) {
	batch.Queue(deleteResourceByIDSQL, resourceID)
}

// DeleteResourceByIDScan implements Querier.DeleteResourceByIDScan.
func (q *DBQuerier) DeleteResourceByIDScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec DeleteResourceByIDBatch: %w", err)
	}
	return cmdTag, err
}

const insertResourceSQL = `INSERT INTO resource(title,
                     description,
                     link)
VALUES(
       $1,
       $2,
       $3
      );`

type InsertResourceParams struct {
	Title       string
	Description string
	Link        string
}

// InsertResource implements Querier.InsertResource.
func (q *DBQuerier) InsertResource(ctx context.Context, params InsertResourceParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "InsertResource")
	cmdTag, err := q.conn.Exec(ctx, insertResourceSQL, params.Title, params.Description, params.Link)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query InsertResource: %w", err)
	}
	return cmdTag, err
}

// InsertResourceBatch implements Querier.InsertResourceBatch.
func (q *DBQuerier) InsertResourceBatch(batch genericBatch, params InsertResourceParams) {
	batch.Queue(insertResourceSQL, params.Title, params.Description, params.Link)
}

// InsertResourceScan implements Querier.InsertResourceScan.
func (q *DBQuerier) InsertResourceScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec InsertResourceBatch: %w", err)
	}
	return cmdTag, err
}

const updateResourceSQL = `UPDATE resource
SET title = $1,
    description = $2,
    link = $3,
    updated_at = NOW()
WHERE resource_id = $4;`

type UpdateResourceParams struct {
	Title       string
	Description string
	Link        string
	ResourceID  int
}

// UpdateResource implements Querier.UpdateResource.
func (q *DBQuerier) UpdateResource(ctx context.Context, params UpdateResourceParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateResource")
	cmdTag, err := q.conn.Exec(ctx, updateResourceSQL, params.Title, params.Description, params.Link, params.ResourceID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateResource: %w", err)
	}
	return cmdTag, err
}

// UpdateResourceBatch implements Querier.UpdateResourceBatch.
func (q *DBQuerier) UpdateResourceBatch(batch genericBatch, params UpdateResourceParams) {
	batch.Queue(updateResourceSQL, params.Title, params.Description, params.Link, params.ResourceID)
}

// UpdateResourceScan implements Querier.UpdateResourceScan.
func (q *DBQuerier) UpdateResourceScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateResourceBatch: %w", err)
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
	return textPreferrer{pgtype.NewValue(t.ValueTranscoder).(pgtype.ValueTranscoder), t.typeName}
}

func (t textPreferrer) TypeName() string {
	return t.typeName
}

// unknownOID means we don't know the OID for a type. This is okay for decoding
// because pgx call DecodeText or DecodeBinary without requiring the OID. For
// encoding parameters, pggen uses textPreferrer if the OID is unknown.
const unknownOID = 0
