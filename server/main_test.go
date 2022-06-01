package main_test

func TestIsBusinessDay(t *testing.T) {
	got := IsBusinessDay("2022-06-01")
	want := true

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
