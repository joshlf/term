package edit

func init() {
	editTests = []editTest{
		{"gofmt.exe", "package main", "package main"},
	}

	// TODO(synful): Make sure this path is actually illegal,
	// not just unlikely to occur in practice.
	illegalPath = "\\illegal\\\n\\path"
}
