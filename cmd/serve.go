package cmd

import "net/http"

func Serve(address string) error {
	fs := http.FileServer(http.Dir("./site"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}