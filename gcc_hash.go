package cpphash

import "errors"

func gcc_std_shift_mix_64bit(v uint64) uint64 {
	return v ^ (v >> 47)
}

func gcc_std_load_8bytes_to_uint64(s string) uint64 {
	if len(s) > 8 {
		panic(errors.New("input not 8 byte"))
	}
	var data uint64
	data = 0
	for i := 0; i < len(s); i++ {
		data = data | (uint64(s[i]) << (i << 3))
	}
	return data
}

func gcc_std_hash_64bit(a string, seed uint64) uint64 {
	const mul = uint64(0xc6a4a793)<<32 + uint64(0x5bd1e995)
	var hash uint64
	hash = seed ^ (uint64(len(a)) * mul)
	for i := 0; (i + 8) < len(a); i = i + 8 {
		data := gcc_std_shift_mix_64bit(gcc_std_load_8bytes_to_uint64(a[i:i+8])*mul) * mul
		hash = hash ^ data
		hash = hash * mul
	}

	if (len(a) & 0x7) != 0 {
		data := gcc_std_load_8bytes_to_uint64(a[len(a)&(^(0x7)):])
		hash = hash ^ data
		hash = hash * mul
	}
	hash = gcc_std_shift_mix_64bit(hash) * mul
	hash = gcc_std_shift_mix_64bit(hash)
	return hash
}

func gcc_std_hash_string(p string) uint64 {
	var seed uint64
	seed = 0xc70f6907
	return gcc_std_hash_64bit(p, seed)
}

func GCCHashString(p string) uint64 {
	return gcc_std_hash_string(p)
}
