session "development" {
    window "editor" {
        run = [
            "nvim"
        ]
    }

    window "server" {
        run = [
            "go run ."
        ]
    }

    window "terminal" {
        send = [
            "echo Hello"
        ]
    }

    focus = "servertestes"
}
