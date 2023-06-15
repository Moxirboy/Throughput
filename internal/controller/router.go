package controller

http.HandleFunc("/kirim", hand.Incoming)
http.HandleFunc("/chiqim", hand.Outcoming)