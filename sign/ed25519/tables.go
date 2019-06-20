package ed25519

import f "github.com/cloudflare/circl/math/fp25519"

var tabSign = [fxV][fx2w1]pointR3{
	{
		pointR3{
			addYX: f.Elt{0x85, 0x3b, 0x8c, 0xf5, 0xc6, 0x93, 0xbc, 0x2f, 0x19, 0x0e, 0x8c, 0xfb, 0xc6, 0x2d, 0x93, 0xcf, 0xc2, 0x42, 0x3d, 0x64, 0x98, 0x48, 0x0b, 0x27, 0x65, 0xba, 0xd4, 0x33, 0x3a, 0x9d, 0xcf, 0x07},
			subYX: f.Elt{0x3e, 0x91, 0x40, 0xd7, 0x05, 0x39, 0x10, 0x9d, 0xb3, 0xbe, 0x40, 0xd1, 0x05, 0x9f, 0x39, 0xfd, 0x09, 0x8a, 0x8f, 0x68, 0x34, 0x84, 0xc1, 0xa5, 0x67, 0x12, 0xf8, 0x98, 0x92, 0x2f, 0xfd, 0x44},
			dt2:   f.Elt{0x68, 0xaa, 0x7a, 0x87, 0x05, 0x12, 0xc9, 0xab, 0x9e, 0xc4, 0xaa, 0xcc, 0x23, 0xe8, 0xd9, 0x26, 0x8c, 0x59, 0x43, 0xdd, 0xcb, 0x7d, 0x1b, 0x5a, 0xa8, 0x65, 0x0c, 0x9f, 0x68, 0x7b, 0x11, 0x6f},
		},
		{
			addYX: f.Elt{0x7c, 0xb0, 0x9e, 0xe6, 0xc5, 0xbf, 0xfa, 0x13, 0x8e, 0x0d, 0x22, 0xde, 0xc8, 0xd1, 0xce, 0x52, 0x02, 0xd5, 0x62, 0x31, 0x71, 0x0e, 0x8e, 0x9d, 0xb0, 0xd6, 0x00, 0xa5, 0x5a, 0x0e, 0xce, 0x72},
			subYX: f.Elt{0x1a, 0x8e, 0x5c, 0xdc, 0xa4, 0xb3, 0x6c, 0x51, 0x18, 0xa0, 0x09, 0x80, 0x9a, 0x46, 0x33, 0xd5, 0xe0, 0x3c, 0x4d, 0x3b, 0xfc, 0x49, 0xa2, 0x43, 0x29, 0xe1, 0x29, 0xa9, 0x93, 0xea, 0x7c, 0x35},
			dt2:   f.Elt{0x08, 0x46, 0x6f, 0x68, 0x7f, 0x0b, 0x7c, 0x9e, 0xad, 0xba, 0x07, 0x61, 0x74, 0x83, 0x2f, 0xfc, 0x26, 0xd6, 0x09, 0xb9, 0x00, 0x34, 0x36, 0x4f, 0x01, 0xf3, 0x48, 0xdb, 0x43, 0xba, 0x04, 0x44},
		},
		{
			addYX: f.Elt{0x4c, 0xda, 0x0d, 0x13, 0x66, 0xfd, 0x82, 0x84, 0x9f, 0x75, 0x5b, 0xa2, 0x17, 0xfe, 0x34, 0xbf, 0x1f, 0xcb, 0xba, 0x90, 0x55, 0x80, 0x83, 0xfd, 0x63, 0xb9, 0x18, 0xf8, 0x5b, 0x5d, 0x94, 0x1e},
			subYX: f.Elt{0xb9, 0xdb, 0x6c, 0x04, 0x88, 0x22, 0xd8, 0x79, 0x83, 0x2f, 0x8d, 0x65, 0x6b, 0xd2, 0xab, 0x1b, 0xdd, 0x65, 0xe5, 0x93, 0x63, 0xf8, 0xa2, 0xd8, 0x3c, 0xf1, 0x4b, 0xc5, 0x99, 0xd1, 0xf2, 0x12},
			dt2:   f.Elt{0x05, 0x4c, 0xb8, 0x3b, 0xfe, 0xf5, 0x9f, 0x2e, 0xd1, 0xb2, 0xb8, 0xff, 0xfe, 0x6d, 0xd9, 0x37, 0xe0, 0xae, 0xb4, 0x5a, 0x51, 0x80, 0x7e, 0x9b, 0x1d, 0xd1, 0x8d, 0x8c, 0x56, 0xb1, 0x84, 0x35},
		},
		{
			addYX: f.Elt{0x39, 0x71, 0x43, 0x34, 0xe3, 0x42, 0x45, 0xa1, 0xf2, 0x68, 0x71, 0xa7, 0xe8, 0x23, 0xfd, 0x9f, 0x86, 0x48, 0xff, 0xe5, 0x96, 0x74, 0xcf, 0x05, 0x49, 0xe2, 0xb3, 0x6c, 0x17, 0x77, 0x2f, 0x6d},
			subYX: f.Elt{0x73, 0x3f, 0xc1, 0xc7, 0x6a, 0x66, 0xa1, 0x20, 0xdd, 0x11, 0xfb, 0x7a, 0x6e, 0xa8, 0x51, 0xb8, 0x3f, 0x9d, 0xa2, 0x97, 0x84, 0xb5, 0xc7, 0x90, 0x7c, 0xab, 0x48, 0xd6, 0x84, 0xa3, 0xd5, 0x1a},
			dt2:   f.Elt{0x63, 0x27, 0x3c, 0x49, 0x4b, 0xfc, 0x22, 0xf2, 0x0b, 0x50, 0xc2, 0x0f, 0xb4, 0x1f, 0x31, 0x0c, 0x2f, 0x53, 0xab, 0xaa, 0x75, 0x6f, 0xe0, 0x69, 0x39, 0x56, 0xe0, 0x3b, 0xb7, 0xa8, 0xbf, 0x45},
		},
	},
	{
		{
			addYX: f.Elt{0x00, 0x45, 0xd9, 0x0d, 0x58, 0x03, 0xfc, 0x29, 0x93, 0xec, 0xbb, 0x6f, 0xa4, 0x7a, 0xd2, 0xec, 0xf8, 0xa7, 0xe2, 0xc2, 0x5f, 0x15, 0x0a, 0x13, 0xd5, 0xa1, 0x06, 0xb7, 0x1a, 0x15, 0x6b, 0x41},
			subYX: f.Elt{0x85, 0x8c, 0xb2, 0x17, 0xd6, 0x3b, 0x0a, 0xd3, 0xea, 0x3b, 0x77, 0x39, 0xb7, 0x77, 0xd3, 0xc5, 0xbf, 0x5c, 0x6a, 0x1e, 0x8c, 0xe7, 0xc6, 0xc6, 0xc4, 0xb7, 0x2a, 0x8b, 0xf7, 0xb8, 0x61, 0x0d},
			dt2:   f.Elt{0xb0, 0x36, 0xc1, 0xe9, 0xef, 0xd7, 0xa8, 0x56, 0x20, 0x4b, 0xe4, 0x58, 0xcd, 0xe5, 0x07, 0xbd, 0xab, 0xe0, 0x57, 0x1b, 0xda, 0x2f, 0xe6, 0xaf, 0xd2, 0xe8, 0x77, 0x42, 0xf7, 0x2a, 0x1a, 0x19},
		},
		{
			addYX: f.Elt{0x6a, 0x6d, 0x6d, 0xd1, 0xfa, 0xf5, 0x03, 0x30, 0xbd, 0x6d, 0xc2, 0xc8, 0xf5, 0x38, 0x80, 0x4f, 0xb2, 0xbe, 0xa1, 0x76, 0x50, 0x1a, 0x73, 0xf2, 0x78, 0x2b, 0x8e, 0x3a, 0x1e, 0x34, 0x47, 0x7b},
			subYX: f.Elt{0xc3, 0x2c, 0x36, 0xdc, 0xc5, 0x45, 0xbc, 0xef, 0x1b, 0x64, 0xd6, 0x65, 0x28, 0xe9, 0xda, 0x84, 0x13, 0xbe, 0x27, 0x8e, 0x3f, 0x98, 0x2a, 0x37, 0xee, 0x78, 0x97, 0xd6, 0xc0, 0x6f, 0xb4, 0x53},
			dt2:   f.Elt{0x58, 0x5d, 0xa7, 0xa3, 0x68, 0xbb, 0x20, 0x30, 0x2e, 0x03, 0xe9, 0xb1, 0xd4, 0x90, 0x72, 0xe3, 0x71, 0xb2, 0x36, 0x3e, 0x73, 0xa0, 0x2e, 0x3d, 0xd1, 0x85, 0x33, 0x62, 0x4e, 0xa7, 0x7b, 0x31},
		},
		{
			addYX: f.Elt{0xbf, 0xc4, 0x38, 0x53, 0xfb, 0x68, 0xa9, 0x77, 0xce, 0x55, 0xf9, 0x05, 0xcb, 0xeb, 0xfb, 0x8c, 0x46, 0xc2, 0x32, 0x7c, 0xf0, 0xdb, 0xd7, 0x2c, 0x62, 0x8e, 0xdd, 0x54, 0x75, 0xcf, 0x3f, 0x33},
			subYX: f.Elt{0x49, 0x50, 0x1f, 0x4e, 0x6e, 0x55, 0x55, 0xde, 0x8c, 0x4e, 0x77, 0x96, 0x38, 0x3b, 0xfe, 0xb6, 0x43, 0x3c, 0x86, 0x69, 0xc2, 0x72, 0x66, 0x1f, 0x6b, 0xf9, 0x87, 0xbc, 0x4f, 0x37, 0x3e, 0x3c},
			dt2:   f.Elt{0xd2, 0x2f, 0x06, 0x6b, 0x08, 0x07, 0x69, 0x77, 0xc0, 0x94, 0xcc, 0xae, 0x43, 0x00, 0x59, 0x6e, 0xa3, 0x63, 0xa8, 0xdd, 0xfa, 0x24, 0x18, 0xd0, 0x35, 0xc7, 0x78, 0xf7, 0x0d, 0xd4, 0x5a, 0x1e},
		},
		{
			addYX: f.Elt{0x45, 0xc1, 0x17, 0x51, 0xf8, 0xed, 0x7e, 0xc7, 0xa9, 0x1a, 0x11, 0x6e, 0x2d, 0xef, 0x0b, 0xd5, 0x3f, 0x98, 0xb0, 0xa3, 0x9d, 0x65, 0xf1, 0xcd, 0x53, 0x4a, 0x8a, 0x18, 0x70, 0x0a, 0x7f, 0x23},
			subYX: f.Elt{0xdd, 0xef, 0xbe, 0x3a, 0x31, 0xe0, 0xbc, 0xbe, 0x6d, 0x5d, 0x79, 0x87, 0xd6, 0xbe, 0x68, 0xe3, 0x59, 0x76, 0x8c, 0x86, 0x0e, 0x7a, 0x92, 0x13, 0x14, 0x8f, 0x67, 0xb3, 0xcb, 0x1a, 0x76, 0x76},
			dt2:   f.Elt{0x56, 0x7a, 0x1c, 0x9d, 0xca, 0x96, 0xf9, 0xf9, 0x03, 0x21, 0xd4, 0xe8, 0xb3, 0xd5, 0xe9, 0x52, 0xc8, 0x54, 0x1e, 0x1b, 0x13, 0xb6, 0xfd, 0x47, 0x7d, 0x02, 0x32, 0x33, 0x27, 0xe2, 0x1f, 0x19},
		},
	},
}

var tabVerif = [1 << (omegaFix - 2)]pointR3{
	{ /* 1P */
		addYX: f.Elt{0x85, 0x3b, 0x8c, 0xf5, 0xc6, 0x93, 0xbc, 0x2f, 0x19, 0x0e, 0x8c, 0xfb, 0xc6, 0x2d, 0x93, 0xcf, 0xc2, 0x42, 0x3d, 0x64, 0x98, 0x48, 0x0b, 0x27, 0x65, 0xba, 0xd4, 0x33, 0x3a, 0x9d, 0xcf, 0x07},
		subYX: f.Elt{0x3e, 0x91, 0x40, 0xd7, 0x05, 0x39, 0x10, 0x9d, 0xb3, 0xbe, 0x40, 0xd1, 0x05, 0x9f, 0x39, 0xfd, 0x09, 0x8a, 0x8f, 0x68, 0x34, 0x84, 0xc1, 0xa5, 0x67, 0x12, 0xf8, 0x98, 0x92, 0x2f, 0xfd, 0x44},
		dt2:   f.Elt{0x68, 0xaa, 0x7a, 0x87, 0x05, 0x12, 0xc9, 0xab, 0x9e, 0xc4, 0xaa, 0xcc, 0x23, 0xe8, 0xd9, 0x26, 0x8c, 0x59, 0x43, 0xdd, 0xcb, 0x7d, 0x1b, 0x5a, 0xa8, 0x65, 0x0c, 0x9f, 0x68, 0x7b, 0x11, 0x6f},
	},
	{ /* 3P */
		addYX: f.Elt{0x30, 0x97, 0xee, 0x4c, 0xa8, 0xb0, 0x25, 0xaf, 0x8a, 0x4b, 0x86, 0xe8, 0x30, 0x84, 0x5a, 0x02, 0x32, 0x67, 0x01, 0x9f, 0x02, 0x50, 0x1b, 0xc1, 0xf4, 0xf8, 0x80, 0x9a, 0x1b, 0x4e, 0x16, 0x7a},
		subYX: f.Elt{0x65, 0xd2, 0xfc, 0xa4, 0xe8, 0x1f, 0x61, 0x56, 0x7d, 0xba, 0xc1, 0xe5, 0xfd, 0x53, 0xd3, 0x3b, 0xbd, 0xd6, 0x4b, 0x21, 0x1a, 0xf3, 0x31, 0x81, 0x62, 0xda, 0x5b, 0x55, 0x87, 0x15, 0xb9, 0x2a},
		dt2:   f.Elt{0x89, 0xd8, 0xd0, 0x0d, 0x3f, 0x93, 0xae, 0x14, 0x62, 0xda, 0x35, 0x1c, 0x22, 0x23, 0x94, 0x58, 0x4c, 0xdb, 0xf2, 0x8c, 0x45, 0xe5, 0x70, 0xd1, 0xc6, 0xb4, 0xb9, 0x12, 0xaf, 0x26, 0x28, 0x5a},
	},
	{ /* 5P */
		addYX: f.Elt{0x33, 0xbb, 0xa5, 0x08, 0x44, 0xbc, 0x12, 0xa2, 0x02, 0xed, 0x5e, 0xc7, 0xc3, 0x48, 0x50, 0x8d, 0x44, 0xec, 0xbf, 0x5a, 0x0c, 0xeb, 0x1b, 0xdd, 0xeb, 0x06, 0xe2, 0x46, 0xf1, 0xcc, 0x45, 0x29},
		subYX: f.Elt{0xba, 0xd6, 0x47, 0xa4, 0xc3, 0x82, 0x91, 0x7f, 0xb7, 0x29, 0x27, 0x4b, 0xd1, 0x14, 0x00, 0xd5, 0x87, 0xa0, 0x64, 0xb8, 0x1c, 0xf1, 0x3c, 0xe3, 0xf3, 0x55, 0x1b, 0xeb, 0x73, 0x7e, 0x4a, 0x15},
		dt2:   f.Elt{0x85, 0x82, 0x2a, 0x81, 0xf1, 0xdb, 0xbb, 0xbc, 0xfc, 0xd1, 0xbd, 0xd0, 0x07, 0x08, 0x0e, 0x27, 0x2d, 0xa7, 0xbd, 0x1b, 0x0b, 0x67, 0x1b, 0xb4, 0x9a, 0xb6, 0x3b, 0x6b, 0x69, 0xbe, 0xaa, 0x43},
	},
	{ /* 7P */
		addYX: f.Elt{0xbf, 0xa3, 0x4e, 0x94, 0xd0, 0x5c, 0x1a, 0x6b, 0xd2, 0xc0, 0x9d, 0xb3, 0x3a, 0x35, 0x70, 0x74, 0x49, 0x2e, 0x54, 0x28, 0x82, 0x52, 0xb2, 0x71, 0x7e, 0x92, 0x3c, 0x28, 0x69, 0xea, 0x1b, 0x46},
		subYX: f.Elt{0xb1, 0x21, 0x32, 0xaa, 0x9a, 0x2c, 0x6f, 0xba, 0xa7, 0x23, 0xba, 0x3b, 0x53, 0x21, 0xa0, 0x6c, 0x3a, 0x2c, 0x19, 0x92, 0x4f, 0x76, 0xea, 0x9d, 0xe0, 0x17, 0x53, 0x2e, 0x5d, 0xdd, 0x6e, 0x1d},
		dt2:   f.Elt{0xa2, 0xb3, 0xb8, 0x01, 0xc8, 0x6d, 0x83, 0xf1, 0x9a, 0xa4, 0x3e, 0x05, 0x47, 0x5f, 0x03, 0xb3, 0xf3, 0xad, 0x77, 0x58, 0xba, 0x41, 0x9c, 0x52, 0xa7, 0x90, 0x0f, 0x6a, 0x1c, 0xbb, 0x9f, 0x7a},
	},
	{ /* 9P */
		addYX: f.Elt{0x2f, 0x63, 0xa8, 0xa6, 0x8a, 0x67, 0x2e, 0x9b, 0xc5, 0x46, 0xbc, 0x51, 0x6f, 0x9e, 0x50, 0xa6, 0xb5, 0xf5, 0x86, 0xc6, 0xc9, 0x33, 0xb2, 0xce, 0x59, 0x7f, 0xdd, 0x8a, 0x33, 0xed, 0xb9, 0x34},
		subYX: f.Elt{0x64, 0x80, 0x9d, 0x03, 0x7e, 0x21, 0x6e, 0xf3, 0x9b, 0x41, 0x20, 0xf5, 0xb6, 0x81, 0xa0, 0x98, 0x44, 0xb0, 0x5e, 0xe7, 0x08, 0xc6, 0xcb, 0x96, 0x8f, 0x9c, 0xdc, 0xfa, 0x51, 0x5a, 0xc0, 0x49},
		dt2:   f.Elt{0x1b, 0xaf, 0x45, 0x90, 0xbf, 0xe8, 0xb4, 0x06, 0x2f, 0xd2, 0x19, 0xa7, 0xe8, 0x83, 0xff, 0xe2, 0x16, 0xcf, 0xd4, 0x93, 0x29, 0xfc, 0xf6, 0xaa, 0x06, 0x8b, 0x00, 0x1b, 0x02, 0x72, 0xc1, 0x73},
	},
	{ /* 11P */
		addYX: f.Elt{0xde, 0x2a, 0x80, 0x8a, 0x84, 0x00, 0xbf, 0x2f, 0x27, 0x2e, 0x30, 0x02, 0xcf, 0xfe, 0xd9, 0xe5, 0x06, 0x34, 0x70, 0x17, 0x71, 0x84, 0x3e, 0x11, 0xaf, 0x8f, 0x6d, 0x54, 0xe2, 0xaa, 0x75, 0x42},
		subYX: f.Elt{0x48, 0x43, 0x86, 0x49, 0x02, 0x5b, 0x5f, 0x31, 0x81, 0x83, 0x08, 0x77, 0x69, 0xb3, 0xd6, 0x3e, 0x95, 0xeb, 0x8d, 0x6a, 0x55, 0x75, 0xa0, 0xa3, 0x7f, 0xc7, 0xd5, 0x29, 0x80, 0x59, 0xab, 0x18},
		dt2:   f.Elt{0xe9, 0x89, 0x60, 0xfd, 0xc5, 0x2c, 0x2b, 0xd8, 0xa4, 0xe4, 0x82, 0x32, 0xa1, 0xb4, 0x1e, 0x03, 0x22, 0x86, 0x1a, 0xb5, 0x99, 0x11, 0x31, 0x44, 0x48, 0xf9, 0x3d, 0xb5, 0x22, 0x55, 0xc6, 0x3d},
	},
	{ /* 13P */
		addYX: f.Elt{0x6d, 0x7f, 0x00, 0xa2, 0x22, 0xc2, 0x70, 0xbf, 0xdb, 0xde, 0xbc, 0xb5, 0x9a, 0xb3, 0x84, 0xbf, 0x07, 0xba, 0x07, 0xfb, 0x12, 0x0e, 0x7a, 0x53, 0x41, 0xf2, 0x46, 0xc3, 0xee, 0xd7, 0x4f, 0x23},
		subYX: f.Elt{0x93, 0xbf, 0x7f, 0x32, 0x3b, 0x01, 0x6f, 0x50, 0x6b, 0x6f, 0x77, 0x9b, 0xc9, 0xeb, 0xfc, 0xae, 0x68, 0x59, 0xad, 0xaa, 0x32, 0xb2, 0x12, 0x9d, 0xa7, 0x24, 0x60, 0x17, 0x2d, 0x88, 0x67, 0x02},
		dt2:   f.Elt{0x78, 0xa3, 0x2e, 0x73, 0x19, 0xa1, 0x60, 0x53, 0x71, 0xd4, 0x8d, 0xdf, 0xb1, 0xe6, 0x37, 0x24, 0x33, 0xe5, 0xa7, 0x91, 0xf8, 0x37, 0xef, 0xa2, 0x63, 0x78, 0x09, 0xaa, 0xfd, 0xa6, 0x7b, 0x49},
	},
	{ /* 15P */
		addYX: f.Elt{0xa0, 0xea, 0xcf, 0x13, 0x03, 0xcc, 0xce, 0x24, 0x6d, 0x24, 0x9c, 0x18, 0x8d, 0xc2, 0x48, 0x86, 0xd0, 0xd4, 0xf2, 0xc1, 0xfa, 0xbd, 0xbd, 0x2d, 0x2b, 0xe7, 0x2d, 0xf1, 0x17, 0x29, 0xe2, 0x61},
		subYX: f.Elt{0x0b, 0xcf, 0x8c, 0x46, 0x86, 0xcd, 0x0b, 0x04, 0xd6, 0x10, 0x99, 0x2a, 0xa4, 0x9b, 0x82, 0xd3, 0x92, 0x51, 0xb2, 0x07, 0x08, 0x30, 0x08, 0x75, 0xbf, 0x5e, 0xd0, 0x18, 0x42, 0xcd, 0xb5, 0x43},
		dt2:   f.Elt{0x16, 0xb5, 0xd0, 0x9b, 0x2f, 0x76, 0x9a, 0x5d, 0xee, 0xde, 0x3f, 0x37, 0x4e, 0xaf, 0x38, 0xeb, 0x70, 0x42, 0xd6, 0x93, 0x7d, 0x5a, 0x2e, 0x03, 0x42, 0xd8, 0xe4, 0x0a, 0x21, 0x61, 0x1d, 0x51},
	},
	{ /* 17P */
		addYX: f.Elt{0x81, 0x9d, 0x0e, 0x95, 0xef, 0x76, 0xc6, 0x92, 0x4f, 0x04, 0xd7, 0xc0, 0xcd, 0x20, 0x46, 0xa5, 0x48, 0x12, 0x8f, 0x6f, 0x64, 0x36, 0x9b, 0xaa, 0xe3, 0x55, 0xb8, 0xdd, 0x24, 0x59, 0x32, 0x6d},
		subYX: f.Elt{0x87, 0xde, 0x20, 0x44, 0x48, 0x86, 0x13, 0x08, 0xb4, 0xed, 0x92, 0xb5, 0x16, 0xf0, 0x1c, 0x8a, 0x25, 0x2d, 0x94, 0x29, 0x27, 0x4e, 0xfa, 0x39, 0x10, 0x28, 0x48, 0xe2, 0x6f, 0xfe, 0xa7, 0x71},
		dt2:   f.Elt{0x54, 0xc8, 0xc8, 0xa5, 0xb8, 0x82, 0x71, 0x6c, 0x03, 0x2a, 0x5f, 0xfe, 0x79, 0x14, 0xfd, 0x33, 0x0c, 0x8d, 0x77, 0x83, 0x18, 0x59, 0xcf, 0x72, 0xa9, 0xea, 0x9e, 0x55, 0xb6, 0xc4, 0x46, 0x47},
	},
	{ /* 19P */
		addYX: f.Elt{0x2b, 0x9a, 0xc6, 0x6d, 0x3c, 0x7b, 0x77, 0xd3, 0x17, 0xf6, 0x89, 0x6f, 0x27, 0xb2, 0xfa, 0xde, 0xb5, 0x16, 0x3a, 0xb5, 0xf7, 0x1c, 0x65, 0x45, 0xb7, 0x9f, 0xfe, 0x34, 0xde, 0x51, 0x9a, 0x5c},
		subYX: f.Elt{0x47, 0x11, 0x74, 0x64, 0xc8, 0x46, 0x85, 0x34, 0x49, 0xc8, 0xfc, 0x0e, 0xdd, 0xae, 0x35, 0x7d, 0x32, 0xa3, 0x72, 0x06, 0x76, 0x9a, 0x93, 0xff, 0xd6, 0xe6, 0xb5, 0x7d, 0x49, 0x63, 0x96, 0x21},
		dt2:   f.Elt{0x67, 0x0e, 0xf1, 0x79, 0xcf, 0xf1, 0x10, 0xf5, 0x5b, 0x51, 0x58, 0xe6, 0xa1, 0xda, 0xdd, 0xff, 0x77, 0x22, 0x14, 0x10, 0x17, 0xa7, 0xc3, 0x09, 0xbb, 0x23, 0x82, 0x60, 0x3c, 0x50, 0x04, 0x48},
	},
	{ /* 21P */
		addYX: f.Elt{0xc7, 0x7f, 0xa3, 0x2c, 0xd0, 0x9e, 0x24, 0xc4, 0xab, 0xac, 0x15, 0xa6, 0xe3, 0xa0, 0x59, 0xa0, 0x23, 0x0e, 0x6e, 0xc9, 0xd7, 0x6e, 0xa9, 0x88, 0x6d, 0x69, 0x50, 0x16, 0xa5, 0x98, 0x33, 0x55},
		subYX: f.Elt{0x75, 0xd1, 0x36, 0x3a, 0xd2, 0x21, 0x68, 0x3b, 0x32, 0x9e, 0x9b, 0xe9, 0xa7, 0x0a, 0xb4, 0xbb, 0x47, 0x8a, 0x83, 0x20, 0xe4, 0x5c, 0x9e, 0x5d, 0x5e, 0x4c, 0xde, 0x58, 0x88, 0x09, 0x1e, 0x77},
		dt2:   f.Elt{0xdf, 0x1e, 0x45, 0x78, 0xd2, 0xf5, 0x12, 0x9a, 0xcb, 0x9c, 0x89, 0x85, 0x79, 0x5d, 0xda, 0x3a, 0x08, 0x95, 0xa5, 0x9f, 0x2d, 0x4a, 0x7f, 0x47, 0x11, 0xa6, 0xf5, 0x8f, 0xd6, 0xd1, 0x5e, 0x5a},
	},
	{ /* 23P */
		addYX: f.Elt{0x83, 0x0e, 0x15, 0xfe, 0x2a, 0x12, 0x95, 0x11, 0xd8, 0x35, 0x4b, 0x7e, 0x25, 0x9a, 0x20, 0xcf, 0x20, 0x1e, 0x71, 0x1e, 0x29, 0xf8, 0x87, 0x73, 0xf0, 0x92, 0xbf, 0xd8, 0x97, 0xb8, 0xac, 0x44},
		subYX: f.Elt{0x59, 0x73, 0x52, 0x58, 0xc5, 0xe0, 0xe5, 0xba, 0x7e, 0x9d, 0xdb, 0xca, 0x19, 0x5c, 0x2e, 0x39, 0xe9, 0xab, 0x1c, 0xda, 0x1e, 0x3c, 0x65, 0x28, 0x44, 0xdc, 0xef, 0x5f, 0x13, 0x60, 0x9b, 0x01},
		dt2:   f.Elt{0x83, 0x4b, 0x13, 0x5e, 0x14, 0x68, 0x60, 0x1e, 0x16, 0x4c, 0x30, 0x24, 0x4f, 0xe6, 0xf5, 0xc4, 0xd7, 0x3e, 0x1a, 0xfc, 0xa8, 0x88, 0x6e, 0x50, 0x92, 0x2f, 0xad, 0xe6, 0xfd, 0x49, 0x0c, 0x15},
	},
	{ /* 25P */
		addYX: f.Elt{0x38, 0x11, 0x47, 0x09, 0x95, 0xf2, 0x7b, 0x8e, 0x51, 0xa6, 0x75, 0x4f, 0x39, 0xef, 0x6f, 0x5d, 0xad, 0x08, 0xa7, 0x25, 0xc4, 0x79, 0xaf, 0x10, 0x22, 0x99, 0xb9, 0x5b, 0x07, 0x5a, 0x2b, 0x6b},
		subYX: f.Elt{0x68, 0xa8, 0xdc, 0x9c, 0x3c, 0x86, 0x49, 0xb8, 0xd0, 0x4a, 0x71, 0xb8, 0xdb, 0x44, 0x3f, 0xc8, 0x8d, 0x16, 0x36, 0x0c, 0x56, 0xe3, 0x3e, 0xfe, 0xc1, 0xfb, 0x05, 0x1e, 0x79, 0xd7, 0xa6, 0x78},
		dt2:   f.Elt{0x76, 0xb9, 0xa0, 0x47, 0x4b, 0x70, 0xbf, 0x58, 0xd5, 0x48, 0x17, 0x74, 0x55, 0xb3, 0x01, 0xa6, 0x90, 0xf5, 0x42, 0xd5, 0xb1, 0x1f, 0x2b, 0xaa, 0x00, 0x5d, 0xd5, 0x4a, 0xfc, 0x7f, 0x5c, 0x72},
	},
	{ /* 27P */
		addYX: f.Elt{0xb2, 0x99, 0xcf, 0xd1, 0x15, 0x67, 0x42, 0xe4, 0x34, 0x0d, 0xa2, 0x02, 0x11, 0xd5, 0x52, 0x73, 0x9f, 0x10, 0x12, 0x8b, 0x7b, 0x15, 0xd1, 0x23, 0xa3, 0xf3, 0xb1, 0x7c, 0x27, 0xc9, 0x4c, 0x79},
		subYX: f.Elt{0xc0, 0x98, 0xd0, 0x1c, 0xf7, 0x2b, 0x80, 0x91, 0x66, 0x63, 0x5e, 0xed, 0xa4, 0x6c, 0x41, 0xfe, 0x4c, 0x99, 0x02, 0x49, 0x71, 0x5d, 0x58, 0xdf, 0xe7, 0xfa, 0x55, 0xf8, 0x25, 0x46, 0xd5, 0x4c},
		dt2:   f.Elt{0x53, 0x50, 0xac, 0xc2, 0x26, 0xc4, 0xf6, 0x4a, 0x58, 0x72, 0xf6, 0x32, 0xad, 0xed, 0x9a, 0xbc, 0x21, 0x10, 0x31, 0x0a, 0xf1, 0x32, 0xd0, 0x2a, 0x85, 0x8e, 0xcc, 0x6f, 0x7b, 0x35, 0x08, 0x70},
	},
	{ /* 29P */
		addYX: f.Elt{0x01, 0x3f, 0x77, 0x38, 0x27, 0x67, 0x88, 0x0b, 0xfb, 0xcc, 0xfb, 0x95, 0xfa, 0xc8, 0xcc, 0xb8, 0xb6, 0x29, 0xad, 0xb9, 0xa3, 0xd5, 0x2d, 0x8d, 0x6a, 0x0f, 0xad, 0x51, 0x98, 0x7e, 0xef, 0x06},
		subYX: f.Elt{0x34, 0x4a, 0x58, 0x82, 0xbb, 0x9f, 0x1b, 0xd0, 0x2b, 0x79, 0xb4, 0xd2, 0x63, 0x64, 0xab, 0x47, 0x02, 0x62, 0x53, 0x48, 0x9c, 0x63, 0x31, 0xb6, 0x28, 0xd4, 0xd6, 0x69, 0x36, 0x2a, 0xa9, 0x13},
		dt2:   f.Elt{0xe5, 0x7d, 0x57, 0xc0, 0x1c, 0x77, 0x93, 0xca, 0x5c, 0xdc, 0x35, 0x50, 0x1e, 0xe4, 0x40, 0x75, 0x71, 0xe0, 0x02, 0xd8, 0x01, 0x0f, 0x68, 0x24, 0x6a, 0xf8, 0x2a, 0x8a, 0xdf, 0x6d, 0x29, 0x3c},
	},
	{ /* 31P */
		addYX: f.Elt{0x13, 0xa7, 0x14, 0xd9, 0xf9, 0x15, 0xad, 0xae, 0x12, 0xf9, 0x8f, 0x8c, 0xf9, 0x7b, 0x2f, 0xa9, 0x30, 0xd7, 0x53, 0x9f, 0x17, 0x23, 0xf8, 0xaf, 0xba, 0x77, 0x0c, 0x49, 0x93, 0xd3, 0x99, 0x7a},
		subYX: f.Elt{0x41, 0x25, 0x1f, 0xbb, 0x2e, 0x4d, 0xeb, 0xfc, 0x1f, 0xb9, 0xad, 0x40, 0xc7, 0x10, 0x95, 0xb8, 0x05, 0xad, 0xa1, 0xd0, 0x7d, 0xa3, 0x71, 0xfc, 0x7b, 0x71, 0x47, 0x07, 0x70, 0x2c, 0x89, 0x0a},
		dt2:   f.Elt{0xe8, 0xa3, 0xbd, 0x36, 0x24, 0xed, 0x52, 0x8f, 0x94, 0x07, 0xe8, 0x57, 0x41, 0xc8, 0xa8, 0x77, 0xe0, 0x9c, 0x2f, 0x26, 0x63, 0x65, 0xa9, 0xa5, 0xd2, 0xf7, 0x02, 0x83, 0xd2, 0x62, 0x67, 0x28},
	},
	{ /* 33P */
		addYX: f.Elt{0x25, 0x5b, 0xe3, 0x3c, 0x09, 0x36, 0x78, 0x4e, 0x97, 0xaa, 0x6b, 0xb2, 0x1d, 0x18, 0xe1, 0x82, 0x3f, 0xb8, 0xc7, 0xcb, 0xd3, 0x92, 0xc1, 0x0c, 0x3a, 0x9d, 0x9d, 0x6a, 0x04, 0xda, 0xf1, 0x32},
		subYX: f.Elt{0xbd, 0xf5, 0x2e, 0xce, 0x2b, 0x8e, 0x55, 0x7c, 0x63, 0xbc, 0x47, 0x67, 0xb4, 0x6c, 0x98, 0xe4, 0xb8, 0x89, 0xbb, 0x3b, 0x9f, 0x17, 0x4a, 0x15, 0x7a, 0x76, 0xf1, 0xd6, 0xa3, 0xf2, 0x86, 0x76},
		dt2:   f.Elt{0x6a, 0x7c, 0x59, 0x6d, 0xa6, 0x12, 0x8d, 0xaa, 0x2b, 0x85, 0xd3, 0x04, 0x03, 0x93, 0x11, 0x8f, 0x22, 0xb0, 0x09, 0xc2, 0x73, 0xdc, 0x91, 0x3f, 0xa6, 0x28, 0xad, 0xa9, 0xf8, 0x05, 0x13, 0x56},
	},
	{ /* 35P */
		addYX: f.Elt{0xd1, 0xae, 0x92, 0xec, 0x8d, 0x97, 0x0c, 0x10, 0xe5, 0x73, 0x6d, 0x4d, 0x43, 0xd5, 0x43, 0xca, 0x48, 0xba, 0x47, 0xd8, 0x22, 0x1b, 0x13, 0x83, 0x2c, 0x4d, 0x5d, 0xe3, 0x53, 0xec, 0xaa},
		subYX: f.Elt{0xd5, 0xc0, 0xb0, 0xe7, 0x28, 0xcc, 0x22, 0x67, 0x53, 0x5c, 0x07, 0xdb, 0xbb, 0xe9, 0x9d, 0x70, 0x61, 0x0a, 0x01, 0xd7, 0xa7, 0x8d, 0xf6, 0xca, 0x6c, 0xcc, 0x57, 0x2c, 0xef, 0x1a, 0x0a, 0x03},
		dt2:   f.Elt{0xaa, 0xd2, 0x3a, 0x00, 0x73, 0xf7, 0xb1, 0x7b, 0x08, 0x66, 0x21, 0x2b, 0x80, 0x29, 0x3f, 0x0b, 0x3e, 0xd2, 0x0e, 0x52, 0x86, 0xdc, 0x21, 0x78, 0x80, 0x54, 0x06, 0x24, 0x1c, 0x9c, 0xbe, 0x20},
	},
	{ /* 37P */
		addYX: f.Elt{0xa6, 0x73, 0x96, 0x24, 0xd8, 0x87, 0x53, 0xe1, 0x93, 0xe4, 0x46, 0xf5, 0x2d, 0xbc, 0x43, 0x59, 0xb5, 0x63, 0x6f, 0xc3, 0x81, 0x9a, 0x7f, 0x1c, 0xde, 0xc1, 0x0a, 0x1f, 0x36, 0xb3, 0x0a, 0x75},
		subYX: f.Elt{0x60, 0x5e, 0x02, 0xe2, 0x4a, 0xe4, 0xe0, 0x20, 0x38, 0xb9, 0xdc, 0xcb, 0x2f, 0x3b, 0x3b, 0xb0, 0x1c, 0x0d, 0x5a, 0xf9, 0x9c, 0x63, 0x5d, 0x10, 0x11, 0xe3, 0x67, 0x50, 0x54, 0x4c, 0x76, 0x69},
		dt2:   f.Elt{0x37, 0x10, 0xf8, 0xa2, 0x83, 0x32, 0x8a, 0x1e, 0xf1, 0xcb, 0x7f, 0xbd, 0x23, 0xda, 0x2e, 0x6f, 0x63, 0x25, 0x2e, 0xac, 0x5b, 0xd1, 0x2f, 0xb7, 0x40, 0x50, 0x07, 0xb7, 0x3f, 0x6b, 0xf9, 0x54},
	},
	{ /* 39P */
		addYX: f.Elt{0x79, 0x92, 0x66, 0x29, 0x04, 0xf2, 0xad, 0x0f, 0x4a, 0x72, 0x7d, 0x7d, 0x04, 0xa2, 0xdd, 0x3a, 0xf1, 0x60, 0x57, 0x8c, 0x82, 0x94, 0x3d, 0x6f, 0x9e, 0x53, 0xb7, 0x2b, 0xc5, 0xe9, 0x7f, 0x3d},
		subYX: f.Elt{0xcd, 0x1e, 0xb1, 0x16, 0xc6, 0xaf, 0x7d, 0x17, 0x79, 0x64, 0x57, 0xfa, 0x9c, 0x4b, 0x76, 0x89, 0x85, 0xe7, 0xec, 0xe6, 0x10, 0xa1, 0xa8, 0xb7, 0xf0, 0xdb, 0x85, 0xbe, 0x9f, 0x83, 0xe6, 0x78},
		dt2:   f.Elt{0x6b, 0x85, 0xb8, 0x37, 0xf7, 0x2d, 0x33, 0x70, 0x8a, 0x17, 0x1a, 0x04, 0x43, 0x5d, 0xd0, 0x75, 0x22, 0x9e, 0xe5, 0xa0, 0x4a, 0xf7, 0x0f, 0x32, 0x42, 0x82, 0x08, 0x50, 0xf3, 0x68, 0xf2, 0x70},
	},
	{ /* 41P */
		addYX: f.Elt{0x47, 0x5f, 0x80, 0xb1, 0x83, 0x45, 0x86, 0x66, 0x19, 0x7c, 0xdd, 0x60, 0xd1, 0xc5, 0x35, 0xf5, 0x06, 0xb0, 0x4c, 0x1e, 0xb7, 0x4e, 0x87, 0xe9, 0xd9, 0x89, 0xd8, 0xfa, 0x5c, 0x34, 0x0d, 0x7c},
		subYX: f.Elt{0x55, 0xf3, 0xdc, 0x70, 0x20, 0x11, 0x24, 0x23, 0x17, 0xe1, 0xfc, 0xe7, 0x7e, 0xc9, 0x0c, 0x38, 0x98, 0xb6, 0x52, 0x35, 0xed, 0xde, 0x1d, 0xb3, 0xb9, 0xc4, 0xb8, 0x39, 0xc0, 0x56, 0x4e, 0x40},
		dt2:   f.Elt{0x8a, 0x33, 0x78, 0x8c, 0x4b, 0x1f, 0x1f, 0x59, 0xe1, 0xb5, 0xe0, 0x67, 0xb1, 0x6a, 0x36, 0xa0, 0x44, 0x3d, 0x5f, 0xb4, 0x52, 0x41, 0xbc, 0x5c, 0x77, 0xc7, 0xae, 0x2a, 0x76, 0x54, 0xd7, 0x20},
	},
	{ /* 43P */
		addYX: f.Elt{0x58, 0xb7, 0x3b, 0xc7, 0x6f, 0xc3, 0x8f, 0x5e, 0x9a, 0xbb, 0x3c, 0x36, 0xa5, 0x43, 0xe5, 0xac, 0x22, 0xc9, 0x3b, 0x90, 0x7d, 0x4a, 0x93, 0xa9, 0x62, 0xec, 0xce, 0xf3, 0x46, 0x1e, 0x8f, 0x2b},
		subYX: f.Elt{0x43, 0xf5, 0xb9, 0x35, 0xb1, 0xfe, 0x74, 0x9d, 0x6c, 0x95, 0x8c, 0xde, 0xf1, 0x7d, 0xb3, 0x84, 0xa9, 0x8b, 0x13, 0x57, 0x07, 0x2b, 0x32, 0xe9, 0xe1, 0x4c, 0x0b, 0x79, 0xa8, 0xad, 0xb8, 0x38},
		dt2:   f.Elt{0x5d, 0xf9, 0x51, 0xdf, 0x9c, 0x4a, 0xc0, 0xb5, 0xac, 0xde, 0x1f, 0xcb, 0xae, 0x52, 0x39, 0x2b, 0xda, 0x66, 0x8b, 0x32, 0x8b, 0x6d, 0x10, 0x1d, 0x53, 0x19, 0xba, 0xce, 0x32, 0xeb, 0x9a, 0x04},
	},
	{ /* 45P */
		addYX: f.Elt{0x31, 0x79, 0xfc, 0x75, 0x0b, 0x7d, 0x50, 0xaa, 0xd3, 0x25, 0x67, 0x7a, 0x4b, 0x92, 0xef, 0x0f, 0x30, 0x39, 0x6b, 0x39, 0x2b, 0x54, 0x82, 0x1d, 0xfc, 0x74, 0xf6, 0x30, 0x75, 0xe1, 0x5e, 0x79},
		subYX: f.Elt{0x7e, 0xfe, 0xdc, 0x63, 0x3c, 0x7d, 0x76, 0xd7, 0x40, 0x6e, 0x85, 0x97, 0x48, 0x59, 0x9c, 0x20, 0x13, 0x7c, 0x4f, 0xe1, 0x61, 0x68, 0x67, 0xb6, 0xfc, 0x25, 0xd6, 0xc8, 0xe0, 0x65, 0xc6, 0x51},
		dt2:   f.Elt{0x81, 0xbd, 0xec, 0x52, 0x0a, 0x5b, 0x4a, 0x25, 0xe7, 0xaf, 0x34, 0xe0, 0x6e, 0x1f, 0x41, 0x5d, 0x31, 0x4a, 0xee, 0xca, 0x0d, 0x4d, 0xa2, 0xe6, 0x77, 0x44, 0xc5, 0x9d, 0xf4, 0x9b, 0xd1, 0x6c},
	},
	{ /* 47P */
		addYX: f.Elt{0x86, 0xc3, 0xaf, 0x65, 0x21, 0x61, 0xfe, 0x1f, 0x10, 0x1b, 0xd5, 0xb8, 0x88, 0x2a, 0x2a, 0x08, 0xaa, 0x0b, 0x99, 0x20, 0x7e, 0x62, 0xf6, 0x76, 0xe7, 0x43, 0x9e, 0x42, 0xa7, 0xb3, 0x01, 0x5e},
		subYX: f.Elt{0xa3, 0x9c, 0x17, 0x52, 0x90, 0x61, 0x87, 0x7e, 0x85, 0x9f, 0x2c, 0x0b, 0x06, 0x0a, 0x1d, 0x57, 0x1e, 0x71, 0x99, 0x84, 0xa8, 0xba, 0xa2, 0x80, 0x38, 0xe6, 0xb2, 0x40, 0xdb, 0xf3, 0x20, 0x75},
		dt2:   f.Elt{0xa1, 0x57, 0x93, 0xd3, 0xe3, 0x0b, 0xb5, 0x3d, 0xa5, 0x94, 0x9e, 0x59, 0xdd, 0x6c, 0x7b, 0x96, 0x6e, 0x1e, 0x31, 0xdf, 0x64, 0x9a, 0x30, 0x1a, 0x86, 0xc9, 0xf3, 0xce, 0x9c, 0x2c, 0x09, 0x71},
	},
	{ /* 49P */
		addYX: f.Elt{0xcf, 0x1d, 0x05, 0x74, 0xac, 0xd8, 0x6b, 0x85, 0x1e, 0xaa, 0xb7, 0x55, 0x08, 0xa4, 0xf6, 0x03, 0xeb, 0x3c, 0x74, 0xc9, 0xcb, 0xe7, 0x4a, 0x3a, 0xde, 0xab, 0x37, 0x71, 0xbb, 0xa5, 0x73, 0x41},
		subYX: f.Elt{0x8c, 0x91, 0x64, 0x03, 0x3f, 0x52, 0xd8, 0x53, 0x1c, 0x6b, 0xab, 0x3f, 0xf4, 0x04, 0xb4, 0xa2, 0xa4, 0xe5, 0x81, 0x66, 0x9e, 0x4a, 0x0b, 0x08, 0xa7, 0x7b, 0x25, 0xd0, 0x03, 0x5b, 0xa1, 0x0e},
		dt2:   f.Elt{0x8a, 0x21, 0xf9, 0xf0, 0x31, 0x6e, 0xc5, 0x17, 0x08, 0x47, 0xfc, 0x1a, 0x2b, 0x6e, 0x69, 0x5a, 0x76, 0xf1, 0xb2, 0xf4, 0x68, 0x16, 0x93, 0xf7, 0x67, 0x3a, 0x4e, 0x4a, 0x61, 0x65, 0xc5, 0x5f},
	},
	{ /* 51P */
		addYX: f.Elt{0x8e, 0x98, 0x90, 0x77, 0xe6, 0xe1, 0x92, 0x48, 0x22, 0xd7, 0x5c, 0x1c, 0x0f, 0x95, 0xd5, 0x01, 0xed, 0x3e, 0x92, 0xe5, 0x9a, 0x81, 0xb0, 0xe3, 0x1b, 0x65, 0x46, 0x9d, 0x40, 0xc7, 0x14, 0x32},
		subYX: f.Elt{0xe5, 0x7a, 0x6d, 0xc4, 0x0d, 0x57, 0x6e, 0x13, 0x8f, 0xdc, 0xf8, 0x54, 0xcc, 0xaa, 0xd0, 0x0f, 0x86, 0xad, 0x0d, 0x31, 0x03, 0x9f, 0x54, 0x59, 0xa1, 0x4a, 0x45, 0x4c, 0x41, 0x1c, 0x71, 0x62},
		dt2:   f.Elt{0x70, 0x17, 0x65, 0x06, 0x74, 0x82, 0x29, 0x13, 0x36, 0x94, 0x27, 0x8a, 0x66, 0xa0, 0xa4, 0x3b, 0x3c, 0x22, 0x5d, 0x18, 0xec, 0xb8, 0xb6, 0xd9, 0x3c, 0x83, 0xcb, 0x3e, 0x07, 0x94, 0xea, 0x5b},
	},
	{ /* 53P */
		addYX: f.Elt{0xf8, 0xd2, 0x43, 0xf3, 0x63, 0xce, 0x70, 0xb4, 0xf1, 0xe8, 0x43, 0x05, 0x8f, 0xba, 0x67, 0x00, 0x6f, 0x7b, 0x11, 0xa2, 0xa1, 0x51, 0xda, 0x35, 0x2f, 0xbd, 0xf1, 0x44, 0x59, 0x78, 0xd0, 0x4a},
		subYX: f.Elt{0xe4, 0x9b, 0xc8, 0x12, 0x09, 0xbf, 0x1d, 0x64, 0x9c, 0x57, 0x6e, 0x7d, 0x31, 0x8b, 0xf3, 0xac, 0x65, 0xb0, 0x97, 0xf6, 0x02, 0x9e, 0xfe, 0xab, 0xec, 0x1e, 0xf6, 0x48, 0xc1, 0xd5, 0xac, 0x3a},
		dt2:   f.Elt{0x01, 0x83, 0x31, 0xc3, 0x34, 0x3b, 0x8e, 0x85, 0x26, 0x68, 0x31, 0x07, 0x47, 0xc0, 0x99, 0xdc, 0x8c, 0xa8, 0x9d, 0xd3, 0x2e, 0x5b, 0x08, 0x34, 0x3d, 0x85, 0x02, 0xd9, 0xb1, 0x0c, 0xff, 0x3a},
	},
	{ /* 55P */
		addYX: f.Elt{0x05, 0x35, 0xc5, 0xf4, 0x0b, 0x43, 0x26, 0x92, 0x83, 0x22, 0x1f, 0x26, 0x13, 0x9c, 0xe4, 0x68, 0xc6, 0x27, 0xd3, 0x8f, 0x78, 0x33, 0xef, 0x09, 0x7f, 0x9e, 0xd9, 0x2b, 0x73, 0x9f, 0xcf, 0x2c},
		subYX: f.Elt{0x5e, 0x40, 0x20, 0x3a, 0xeb, 0xc7, 0xc5, 0x87, 0xc9, 0x56, 0xad, 0xed, 0xef, 0x11, 0xe3, 0x8e, 0xf9, 0xd5, 0x29, 0xad, 0x48, 0x2e, 0x25, 0x29, 0x1d, 0x25, 0xcd, 0xf4, 0x86, 0x7e, 0x0e, 0x11},
		dt2:   f.Elt{0xe4, 0xf5, 0x03, 0xd6, 0x9e, 0xd8, 0xc0, 0x57, 0x0c, 0x20, 0xb0, 0xf0, 0x28, 0x86, 0x88, 0x12, 0xb7, 0x3b, 0x2e, 0xa0, 0x09, 0x27, 0x17, 0x53, 0x37, 0x3a, 0x69, 0xb9, 0xe0, 0x57, 0xc5, 0x05},
	},
	{ /* 57P */
		addYX: f.Elt{0xb0, 0x0e, 0xc2, 0x89, 0xb0, 0xbb, 0x76, 0xf7, 0x5c, 0xd8, 0x0f, 0xfa, 0xf6, 0x5b, 0xf8, 0x61, 0xfb, 0x21, 0x44, 0x63, 0x4e, 0x3f, 0xb9, 0xb6, 0x05, 0x12, 0x86, 0x41, 0x08, 0xef, 0x9f, 0x28},
		subYX: f.Elt{0x6f, 0x7e, 0xc9, 0x1f, 0x31, 0xce, 0xf9, 0xd8, 0xae, 0xfd, 0xf9, 0x11, 0x30, 0x26, 0x3f, 0x7a, 0xdd, 0x25, 0xed, 0x8b, 0xa0, 0x7e, 0x5b, 0xe1, 0x5a, 0x87, 0xe9, 0x8f, 0x17, 0x4c, 0x15, 0x6e},
		dt2:   f.Elt{0xbf, 0x9a, 0xd6, 0xfe, 0x36, 0x63, 0x61, 0xcf, 0x4f, 0xc9, 0x35, 0x83, 0xe7, 0xe4, 0x16, 0x9b, 0xe7, 0x7f, 0x3a, 0x75, 0x65, 0x97, 0x78, 0x13, 0x19, 0xa3, 0x5c, 0xa9, 0x42, 0xf6, 0xfb, 0x6a},
	},
	{ /* 59P */
		addYX: f.Elt{0xcc, 0xa8, 0x13, 0xf9, 0x70, 0x50, 0xe5, 0x5d, 0x61, 0xf5, 0x0c, 0x2b, 0x7b, 0x16, 0x1d, 0x7d, 0x89, 0xd4, 0xea, 0x90, 0xb6, 0x56, 0x29, 0xda, 0xd9, 0x1e, 0x80, 0xdb, 0xce, 0x93, 0xc0, 0x12},
		subYX: f.Elt{0xc1, 0xd2, 0xf5, 0x62, 0x0c, 0xde, 0xa8, 0x7d, 0x9a, 0x7b, 0x0e, 0xb0, 0xa4, 0x3d, 0xfc, 0x98, 0xe0, 0x70, 0xad, 0x0d, 0xda, 0x6a, 0xeb, 0x7d, 0xc4, 0x38, 0x50, 0xb9, 0x51, 0xb8, 0xb4, 0x0d},
		dt2:   f.Elt{0x0f, 0x19, 0xb8, 0x08, 0x93, 0x7f, 0x14, 0xfc, 0x10, 0xe3, 0x1a, 0xa1, 0xa0, 0x9d, 0x96, 0x06, 0xfd, 0xd7, 0xc7, 0xda, 0x72, 0x55, 0xe7, 0xce, 0xe6, 0x5c, 0x63, 0xc6, 0x99, 0x87, 0xaa, 0x33},
	},
	{ /* 61P */
		addYX: f.Elt{0xb1, 0x6c, 0x15, 0xfc, 0x88, 0xf5, 0x48, 0x83, 0x27, 0x6d, 0x0a, 0x1a, 0x9b, 0xba, 0xa2, 0x6d, 0xb6, 0x5a, 0xca, 0x87, 0x5c, 0x2d, 0x26, 0xe2, 0xa6, 0x89, 0xd5, 0xc8, 0xc1, 0xd0, 0x2c, 0x21},
		subYX: f.Elt{0xf2, 0x5c, 0x08, 0xbd, 0x1e, 0xf5, 0x0f, 0xaf, 0x1f, 0x3f, 0xd3, 0x67, 0x89, 0x1a, 0xf5, 0x78, 0x3c, 0x03, 0x60, 0x50, 0xe1, 0xbf, 0xc2, 0x6e, 0x86, 0x1a, 0xe2, 0xe8, 0x29, 0x6f, 0x3c, 0x23},
		dt2:   f.Elt{0x81, 0xc7, 0x18, 0x7f, 0x10, 0xd5, 0xf4, 0xd2, 0x28, 0x9d, 0x7e, 0x52, 0xf2, 0xcd, 0x2e, 0x12, 0x41, 0x33, 0x3d, 0x3d, 0x2a, 0x86, 0x0a, 0xa7, 0xe3, 0x4c, 0x91, 0x11, 0x89, 0x77, 0xb7, 0x1d},
	},
	{ /* 63P */
		addYX: f.Elt{0xb6, 0x1a, 0x70, 0xdd, 0x69, 0x47, 0x39, 0xb3, 0xa5, 0x8d, 0xcf, 0x19, 0xd4, 0xde, 0xb8, 0xe2, 0x52, 0xc8, 0x2a, 0xfd, 0x61, 0x41, 0xdf, 0x15, 0xbe, 0x24, 0x7d, 0x01, 0x8a, 0xca, 0xe2, 0x7a},
		subYX: f.Elt{0x6f, 0xc2, 0x6b, 0x7c, 0x39, 0x52, 0xf3, 0xdd, 0x13, 0x01, 0xd5, 0x53, 0xcc, 0xe2, 0x97, 0x7a, 0x30, 0xa3, 0x79, 0xbf, 0x3a, 0xf4, 0x74, 0x7c, 0xfc, 0xad, 0xe2, 0x26, 0xad, 0x97, 0xad, 0x31},
		dt2:   f.Elt{0x62, 0xb9, 0x20, 0x09, 0xed, 0x17, 0xe8, 0xb7, 0x9d, 0xda, 0x19, 0x3f, 0xcc, 0x18, 0x85, 0x1e, 0x64, 0x0a, 0x56, 0x25, 0x4f, 0xc1, 0x91, 0xe4, 0x83, 0x2c, 0x62, 0xa6, 0x53, 0xfc, 0xd1, 0x1e},
	},
}