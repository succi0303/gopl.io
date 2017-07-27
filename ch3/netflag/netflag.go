func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FalgBroadcast }
func IsCast(v Flags) bool { return v&(FlagBroadcast | FlagMulticast) != 0 }

func main() {
    var v Flags = FlagMulticase | FlagUp
    fmt.Printf("%b %t\n", v, IsUp(v))
    TrunDown(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    SetBroadcast(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    fmt.Printf("%b %t\n", v, IsCast(v))
}
