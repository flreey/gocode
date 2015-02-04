fork from: https://github.com/nsf/gocode

modify returns for vim plugin

previous:
     `<c-x><c-o> fmt.Errorf(`

now:
    `<c-x><c-o> fmt.Errorf(<format string>, <a ...interface{}>)`
