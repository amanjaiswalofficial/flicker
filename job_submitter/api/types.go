package api

type Source struct {
    Name string
    URL  string
}

type Task struct {
    Name string
    Args []string
}

type Job struct {
    Source Source
    Tasks  []Task
}
