fork from: https://github.com/nsf/gocode

modify returns for vim plugin

previous:
     `<c-x><c-o> fmt.Errorf(`

now:
    `<c-x><c-o> fmt.Errorf(<format string>, <a ...interface{}>)`

withing vim plugin: https://github.com/flreey/switch_region, press <tab> will jump to params
