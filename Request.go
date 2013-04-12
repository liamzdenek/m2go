package main

type Request struct {
    sender string;
    conn_id string;
    path string;
    body string;
    headers []Header; 
}
