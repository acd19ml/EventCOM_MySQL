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

	QueryFormSQL = `
	SELECT 
    h.id,
    h.name

    FROM 
    	Head h
	`

	DescribeFormSQL = `
	SELECT 
    h.id,
    h.name, 
    f.label, 
    f.type, 
    f.required, 
    f.description, 
    f.min_value, 
    f.max_value, 
    f.min_date, 
    f.max_date, 
    f.multiple_selection, 
    f.options
    FROM 
    	Head h
    	RIGHT JOIN  Field f ON h.id = f.head_id
	`
)
