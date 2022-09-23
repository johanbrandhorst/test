// Code generated via scripts/generate.sh. DO NOT EDIT.

package must

import (
	"io/fs"
	"regexp"

	"github.com/shoenig/test/interfaces"
	"github.com/shoenig/test/internal/assertions"
	"github.com/shoenig/test/internal/constraints"
)

// Nil asserts a is nil.
func Nil(t T, a any, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Nil(a), scripts...)
}

// NotNil asserts a is not nil.
func NotNil(t T, a any, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotNil(a), scripts...)
}

// True asserts that condition is true.
func True(t T, condition bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.True(condition), scripts...)
}

// False asserts condition is false.
func False(t T, condition bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.False(condition), scripts...)
}

// Unreachable asserts a code path is not executed.
func Unreachable(t T, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Unreachable(), scripts...)
}

// Error asserts err is a non-nil error.
func Error(t T, err error, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Error(err), scripts...)
}

// EqError asserts err contains message msg.
func EqError(t T, err error, msg string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.EqError(err, msg), scripts...)
}

// ErrorIs asserts err
func ErrorIs(t T, err error, target error, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.ErrorIs(err, target), scripts...)
}

// NoError asserts err is a nil error.
func NoError(t T, err error, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NoError(err), scripts...)
}

// Eq asserts a and b are equal using cmp.Equal.
func Eq[A any](t T, a, b A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Eq(a, b), scripts...)
}

// EqOp asserts a == b.
func EqOp[C comparable](t T, a, b C, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.EqOp(a, b), scripts...)
}

// EqFunc asserts a and b are equal using eq.
func EqFunc[A any](t T, a, b A, eq func(a, b A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.EqFunc(a, b, eq), scripts...)
}

// NotEq asserts a and b are not equal using cmp.Equal.
func NotEq[A any](t T, a, b A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotEq(a, b), scripts...)
}

// NotEqOp asserts a != b.
func NotEqOp[C comparable](t T, a, b C, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotEqOp(a, b), scripts...)
}

// NotEqFunc asserts a and b are not equal using eq.
func NotEqFunc[A any](t T, a, b A, eq func(a, b A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotEqFunc(a, b, eq), scripts...)
}

// EqJSON asserts a and b are equivalent JSON.
func EqJSON(t T, a, b string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.EqJSON(a, b), scripts...)
}

// SliceEqFunc asserts elements of a and b are the same using eq.
func SliceEqFunc[A any](t T, a, b []A, eq func(a, b A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.EqSliceFunc(a, b, eq), scripts...)
}

// Equal asserts a.Equal(b).
func Equal[E interfaces.EqualFunc[E]](t T, a, b E, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Equal(a, b), scripts...)
}

// NotEqual asserts !a.Equal(b).
func NotEqual[E interfaces.EqualFunc[E]](t T, a, b E, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotEqual(a, b), scripts...)
}

// Lesser asserts a.Less(b).
func Lesser[L interfaces.LessFunc[L]](t T, a, b L, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Lesser(a, b), scripts...)
}

// SliceEqual asserts a[n].Equal(b[n]) for each element n in slices a and b.
func SliceEqual[E interfaces.EqualFunc[E]](t T, a, b []E, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceEqual(a, b), scripts...)
}

// SliceEmpty asserts slice is empty.
func SliceEmpty[A any](t T, slice []A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceEmpty(slice), scripts...)
}

// SliceNotEmpty asserts slice is not empty.
func SliceNotEmpty[A any](t T, slice []A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceNotEmpty(slice), scripts...)
}

// SliceLen asserts slice is of length n.
func SliceLen[A any](t T, n int, slice []A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceLen(n, slice), scripts...)
}

// Len asserts slice is of length n.
//
// Shorthand function for SliceLen. For checking Len() of a struct,
// use the Length() assertion.
func Len[A any](t T, n int, slice []A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceLen(n, slice), scripts...)
}

// SliceContainsOp asserts item exists in slice using == operator.
func SliceContainsOp[C comparable](t T, slice []C, item C, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceContainsOp(slice, item), scripts...)
}

// SliceContainsFunc asserts item exists in slice, using eq to compare elements.
func SliceContainsFunc[A any](t T, slice []A, item A, eq func(a, b A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceContainsFunc(slice, item, eq), scripts...)
}

// SliceContainsEqual asserts item exists in slice, using Equal to compare elements.
func SliceContainsEqual[E interfaces.EqualFunc[E]](t T, slice []E, item E, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceContainsEqual(slice, item), scripts...)
}

// SliceContains asserts item exists in slice, using cmp.Equal to compare elements.
func SliceContains[A any](t T, slice []A, item A, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.SliceContains(slice, item), scripts...)
}

// Positive asserts n > 0.
func Positive[N interfaces.Number](t T, n N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Positive(n), scripts...)
}

// Negative asserts n < 0.
func Negative[N interfaces.Number](t T, n N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Negative(n), scripts...)
}

// Zero asserts n == 0.
func Zero[N interfaces.Number](t T, n N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Zero(n), scripts...)
}

// NonZero asserts n != 0.
func NonZero[N interfaces.Number](t T, n N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NonZero(n), scripts...)
}

// One asserts n == 1.
func One[N interfaces.Number](t T, n N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.One(n), scripts...)
}

// Less asserts a < b.
func Less[O constraints.Ordered](t T, a, b O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Less(a, b), scripts...)
}

// LessEq asserts a <= b.
func LessEq[O constraints.Ordered](t T, a, b O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.LessEq(a, b), scripts...)
}

// Greater asserts a > b.
func Greater[O constraints.Ordered](t T, a, b O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Greater(a, b), scripts...)
}

// GreaterEq asserts a >= b.
func GreaterEq[O constraints.Ordered](t T, a, b O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.GreaterEq(a, b), scripts...)
}

// Ascending asserts slice[n] <= slice[n+1] for each element n.
func Ascending[O constraints.Ordered](t T, slice []O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Ascending(slice), scripts...)
}

// AscendingFunc asserts slice[n] is less than slice[n+1] for each element n using the less comparator.
func AscendingFunc[A any](t T, slice []A, less func(A, A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.AscendingFunc(slice, less), scripts...)
}

// AscendingLess asserts slice[n].Less(slice[n+1]) for each element n.
func AscendingLess[L interfaces.LessFunc[L]](t T, slice []L, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.AscendingLess(slice), scripts...)
}

// Descending asserts slice[n] >= slice[n+1] for each element n.
func Descending[O constraints.Ordered](t T, slice []O, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Descending(slice), scripts...)
}

// DescendingFunc asserts slice[n+1] is less than slice[n] for each element n using the less comparator.
func DescendingFunc[A any](t T, slice []A, less func(A, A) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.DescendingFunc(slice, less), scripts...)
}

// DescendingLess asserts slice[n+1].Less(slice[n]) for each element n.
func DescendingLess[L interfaces.LessFunc[L]](t T, slice []L, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.DescendingLess(slice), scripts...)
}

// InDelta asserts a and b are within delta of each other.
func InDelta[N interfaces.Number](t T, a, b, delta N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.InDelta(a, b, delta), scripts...)
}

// InDeltaSlice asserts each element a[n] is within delta of b[n].
func InDeltaSlice[N interfaces.Number](t T, a, b []N, delta N, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.InDeltaSlice(a, b, delta), scripts...)
}

// MapEq asserts maps a and b contain the same key/value pairs, using
// cmp.Equal function to compare values.
func MapEq[M1, M2 interfaces.Map[K, V], K comparable, V any](t T, a M1, b M2, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapEq[M1, M2, K, V](a, b), scripts...)
}

// MapEqFunc asserts maps a and b contain the same key/value pairs, using eq to
// compare values.
func MapEqFunc[M1, M2 interfaces.Map[K, V], K comparable, V any](t T, a M1, b M2, eq func(V, V) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapEqFunc[M1, M2, K, V](a, b, eq), scripts...)
}

// MapEqual asserts maps a and b contain the same key/value pairs, using Equal
// method to compare values
func MapEqual[M interfaces.MapEqualFunc[K, V], K comparable, V interfaces.EqualFunc[V]](t T, a, b M, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapEqual[M, K, V](a, b), scripts...)
}

// MapLen asserts map is of size n.
func MapLen[M ~map[K]V, K comparable, V any](t T, n int, m M, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapLen[M, K, V](n, m), scripts...)
}

// MapEmpty asserts map is empty.
func MapEmpty[M ~map[K]V, K comparable, V any](t T, m M, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapEmpty[M, K, V](m), scripts...)
}

// MapContainsKeys asserts m contains each key in keys.
func MapContainsKeys[M ~map[K]V, K comparable, V any](t T, m M, keys []K, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapContainsKeys[M, K, V](m, keys), scripts...)
}

// MapContainsValues asserts m contains each value in values.
func MapContainsValues[M ~map[K]V, K comparable, V any](t T, m M, values []V, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapContainsValues[M, K, V](m, values), scripts...)
}

// MapContainsValuesFunc asserts m contains each value in values using the eq function.
func MapContainsValuesFunc[M ~map[K]V, K comparable, V any](t T, m M, values []V, eq func(V, V) bool, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapContainsValuesFunc[M, K, V](m, values, eq), scripts...)
}

func MapContainsValuesEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](t T, m M, values []V, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.MapContainsValuesEqual[M, K, V](m, values), scripts...)
}

// FileExists asserts file exists on system.
//
// Often os.DirFS is used to interact with the the host filesystem.
// Example,
// FileExists(t, os.DirFS("/etc"), "hosts")
func FileExists(t T, system fs.FS, file string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.FileExists(system, file), scripts...)
}

// FileNotExists asserts file does not exist on system.
//
// Often os.DirFS is used to interact with the host filesystem.
// Example,
// FileNotExist(t, os.DirFS("/bin"), "exploit.exe")
func FileNotExists(t T, system fs.FS, file string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.FileNotExists(system, file), scripts...)
}

// DirExists asserts directory exists on system.
//
// Often os.DirFS is used to interact with the host filesystem.
// Example,
// DirExists(t, os.DirFS("/usr/local"), "bin")
func DirExists(t T, system fs.FS, directory string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.DirExists(system, directory), scripts...)
}

// DirNotExists asserts directory does not exist on system.
//
// Often os.DirFS is used to interact with the host filesystem.
// Example,
// DirNotExists(t, os.DirFS("/tmp"), "scratch")
func DirNotExists(t T, system fs.FS, directory string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.DirNotExists(system, directory), scripts...)
}

// FileMode asserts the file or directory at path has exactly
// the given permission bits.
//
// Often os.DirFS is used to interact with the host filesystem.
// Example,
// FileMode(t, os.DirFS("/bin"), "find", 0655)
func FileMode(t T, system fs.FS, path string, permissions fs.FileMode, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.FileMode(system, path, permissions), scripts...)
}

// FileContains asserts the file contains content as a substring.
//
// Often os.DirFS is used to interact with the host filesystem.
// Example,
// FileContains(t, os.DirFS("/etc"), "hosts", "localhost")
func FileContains(t T, system fs.FS, file, content string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.FileContains(system, file, content), scripts...)
}

// FilePathValid asserts path is a valid file path.
func FilePathValid(t T, path string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.FilePathValid(path), scripts...)
}

// StrEqFold asserts first and second are equivalent, ignoring case.
func StrEqFold(t T, first, second string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrEqFold(first, second), scripts...)
}

// StrNotEqFold asserts first and second are not equivalent, ignoring case.
func StrNotEqFold(t T, first, second string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotEqFold(first, second), scripts...)
}

// StrContains asserts s contains substring sub.
func StrContains(t T, s, sub string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrContains(s, sub), scripts...)
}

// StrContainsFold asserts s contains substring sub, ignoring case.
func StrContainsFold(t T, s, sub string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrContainsFold(s, sub), scripts...)
}

// StrNotContains asserts s does not contain substring sub.
func StrNotContains(t T, s, sub string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotContains(s, sub), scripts...)
}

// StrNotContainsFold asserts s does not contain substring sub, ignoring case.
func StrNotContainsFold(t T, s, sub string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotContainsFold(s, sub), scripts...)
}

// StrContainsAny asserts s contains at least one character in chars.
func StrContainsAny(t T, s, chars string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrContainsAny(s, chars), scripts...)
}

// StrNotContainsAny asserts s does not contain any character in chars.
func StrNotContainsAny(t T, s, chars string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotContainsAny(s, chars), scripts...)
}

// StrCount asserts s contains exactly count instances of substring sub.
func StrCount(t T, s, sub string, count int, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrCount(s, sub, count), scripts...)
}

// StrContainsFields asserts that fields is a subset of the result of strings.Fields(s).
func StrContainsFields(t T, s string, fields []string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrContainsFields(s, fields), scripts...)
}

// StrHasPrefix asserts that s starts with prefix.
func StrHasPrefix(t T, s, prefix string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrHasPrefix(s, prefix), scripts...)
}

// StrNotHasPrefix asserts that s does not start with prefix.
func StrNotHasPrefix(t T, s, prefix string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotHasPrefix(s, prefix), scripts...)
}

// StrHasSuffix asserts that s ends with suffix.
func StrHasSuffix(t T, s, suffix string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrHasSuffix(s, suffix), scripts...)
}

// StrNotHasSuffix asserts that s does not end with suffix.
func StrNotHasSuffix(t T, s, suffix string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.StrNotHasSuffix(s, suffix), scripts...)
}

// RegexMatch asserts regular expression re matches string s.
func RegexMatch(t T, re *regexp.Regexp, s string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.RegexMatch(re, s), scripts...)
}

// RegexCompiles asserts expr compiles as a valid regular expression.
func RegexCompiles(t T, expr string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.RegexpCompiles(expr), scripts...)
}

// RegexCompilesPOSIX asserts expr compiles as a valid POSIX regular expression.
func RegexCompilesPOSIX(t T, expr string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.RegexpCompilesPOSIX(expr), scripts...)
}

// UUIDv4 asserts id meets the criteria of a v4 UUID.
func UUIDv4(t T, id string, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.UUIDv4(id), scripts...)
}

// Size asserts s.Size() is equal to exp.
func Size(t T, exp int, s interfaces.SizeFunc, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Size(exp, s), scripts...)
}

// Length asserts l.Len() is equal to exp.
func Length(t T, exp int, l interfaces.LengthFunc, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Length(exp, l), scripts...)
}

// Empty asserts e.Empty() is true.
func Empty(t T, e interfaces.EmptyFunc, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Empty(e), scripts...)
}

// NotEmpty asserts e.Empty() is false.
func NotEmpty(t T, e interfaces.EmptyFunc, scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotEmpty(e), scripts...)
}

// Contains asserts container.Contains(element) is true.
func Contains[C any](t T, element C, container interfaces.Contains[C], scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.Contains(element, container), scripts...)
}

// NotContains asserts container.Contains(element) is false.
func NotContains[C any](t T, element C, container interfaces.Contains[C], scripts ...PostScript) {
	t.Helper()
	invoke(t, assertions.NotContains(element, container), scripts...)
}
