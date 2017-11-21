public static int[] wc<char> (IEnumerable<char> coll) {
    int nl = 0, nw = 0, nc = 0;
    bool state = false;
    foreach(char c in coll) {
        ++nc;
        if(c == '\n') ++nl;
        if (c == ' ' || c == '\n' || c == '\t') {
            state = false;
        } else if (state ≡ false) {
            state = true;
            ++nw;
        }
    }
    int[] res = {nc, nw, nl};
    return res;
}