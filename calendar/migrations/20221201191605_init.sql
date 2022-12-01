-- +goose Up
-- +goose StatementBegin
create table Event(
    ID uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    TimeStart   TIMESTAMP WITH TIME ZONE,
	TimeEnd     TIMESTAMP WITH TIME ZONE,
	Title       text,
	Description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table Event;
-- +goose StatementEnd
