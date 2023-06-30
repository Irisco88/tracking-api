package postgres

import (
	"context"
	sqlmaker "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type FMSDataBaseConn interface {
	GetPgConn() *pgxpool.Pool
	GetSQLBuilder() sqlmaker.StatementBuilderType
}

var _ FMSDataBaseConn = &FMSDataBase{}

type FMSDataBase struct {
	postgresConn  *pgxpool.Pool
	selectBuilder sqlmaker.StatementBuilderType
}

func (adb *FMSDataBase) GetPgConn() *pgxpool.Pool {
	return adb.postgresConn
}

func (adb *FMSDataBase) GetSQLBuilder() sqlmaker.StatementBuilderType {
	adb.selectBuilder = sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar)
	return adb.selectBuilder
}

func NewFmsDataBase(rawConn *pgxpool.Pool) *FMSDataBase {
	return &FMSDataBase{
		selectBuilder: sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar),
		postgresConn:  rawConn,
	}
}

func ConnectToFmsDB(databaseURL string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rawConn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return rawConn, nil
}
