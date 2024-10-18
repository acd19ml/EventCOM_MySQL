package impl

const (
	InsertHeadSQL = `
	INSERT INTO Head (
	id, 
	name, 
	created_at, 
	updated_at
	) 
	VALUES 
			(?, ?, ?, ?);

	`

	InsertFieldSQL = `
	INSERT INTO Field (
	id,
	head_id, 
	label, 
	type, 
	required, 
	description, 
	min_value, 
	max_value, 
	min_date, 
	max_date, 
	multiple_selection, 
	options
	)
	VALUES 
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

	`
)
