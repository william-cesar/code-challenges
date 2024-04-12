func mergeAlternately(w1 string, w2 string) string {
    var concat string 

    for len(w1) > 0 && len(w2) > 0 {
        concat += string(w1[0]) + string(w2[0])
        w1 = w1[1:]
        w2 = w2[1:]
    }

    if len(w1) > 0 {
        concat += w1
    }

    if len(w2) > 0 {
        concat += w2
    }

    return concat
}
