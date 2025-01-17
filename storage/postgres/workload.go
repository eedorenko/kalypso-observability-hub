package postgres

import "database/sql"

type Workload struct {
	Id                int
	Name              string
	Description       string
	SourceStorageType string
	SourceEndpoint    string
	ApplicationId     int
}

// make sure that the Workload implements the Entity interface
var _ Entity = (*Workload)(nil)

func (wrk *Workload) update(conn *sql.DB) (Entity, error) {
	var id int32
	err := conn.QueryRow(`INSERT INTO workload (name, description, source_storage_type, source_endpoint, application_id) VALUES ($1, $2, $3, $4, $5)
						 ON CONFLICT (application_id, name) DO
						 UPDATE SET description=$2,
						 			source_storage_type=$3,
									source_endpoint=$4,
								    updated_on=current_timestamp,
						            updated_by=current_user
						 RETURNING id`, wrk.Name, wrk.Description, wrk.SourceStorageType, wrk.SourceEndpoint, wrk.ApplicationId).Scan(&id)
	if err != nil {
		return nil, err
	}
	wrk.Id = int(id)
	return wrk, nil
}
