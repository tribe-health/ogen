// Code generated by ogen, DO NOT EDIT.

package api

// setDefaults set default value of fields.
func (s *Drive) setDefaults() {
	{
		val := string("Unsafe")

		s.CacheType.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Logger) setDefaults() {
	{
		val := LoggerLevel("Warning")

		s.Level.SetTo(val)
	}
	{
		val := bool(false)

		s.ShowLevel.SetTo(val)
	}
	{
		val := bool(false)

		s.ShowLogOrigin.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *MmdsConfig) setDefaults() {
	{
		val := string("169.254.169.254")

		s.Ipv4Address.SetTo(val)
	}
}
