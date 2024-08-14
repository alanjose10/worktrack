


Tasks:
    ID          int primary key
    Project     string
    Name        string
    Status      int
    CreatedOn   date
    UpdatedOn   date

Notes:
    ID int      primary key
    TaskID      int foreign key to Tasks.ID
    Content     string
    AddedOn     date

