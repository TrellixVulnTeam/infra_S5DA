// Code generated by cproto. DO NOT EDIT.

package adminpb

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"weetbix.internal.admin.Admin",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 236, 123, 77, 140, 27, 199,
			182, 30, 155, 205, 25, 143, 202, 250, 109, 233, 218, 186, 180, 174,
			124, 68, 123, 44, 210, 225, 52, 231, 71, 191, 163, 171, 139, 203,
			225, 244, 72, 173, 59, 34, 231, 146, 156, 145, 37, 195, 150, 138,
			221, 69, 178, 173, 102, 23, 221, 85, 28, 138, 86, 4, 60, 4,
			8, 30, 240, 30, 240, 144, 4, 8, 144, 44, 146, 0, 201, 42,
			15, 1, 178, 8, 144, 213, 3, 30, 2, 36, 200, 230, 1, 217,
			101, 145, 69, 118, 217, 38, 155, 44, 2, 36, 8, 130, 58, 213,
			77, 114, 164, 145, 100, 223, 4, 1, 30, 224, 129, 198, 238, 170,
			174, 58, 117, 206, 169, 83, 231, 124, 231, 84, 15, 249, 75, 139,
			124, 210, 227, 188, 23, 178, 202, 48, 230, 146, 119, 70, 221, 10,
			27, 12, 229, 196, 198, 166, 117, 70, 191, 180, 211, 151, 133, 15,
			200, 130, 163, 222, 111, 189, 34, 231, 61, 62, 176, 95, 123, 191,
			69, 240, 237, 158, 106, 238, 25, 79, 210, 215, 61, 30, 210, 168,
			103, 243, 184, 55, 91, 70, 78, 134, 76, 84, 158, 71, 124, 28,
			233, 37, 135, 157, 255, 97, 24, 255, 52, 107, 222, 219, 219, 250,
			243, 236, 229, 123, 122, 230, 94, 50, 220, 126, 196, 194, 240, 119,
			106, 112, 91, 205, 123, 240, 191, 207, 146, 69, 43, 119, 57, 179,
			113, 150, 252, 213, 73, 98, 156, 180, 204, 203, 25, 107, 253, 223,
			156, 4, 156, 224, 241, 16, 182, 70, 221, 46, 139, 5, 172, 128,
			38, 117, 85, 128, 79, 37, 133, 32, 146, 44, 246, 250, 52, 234,
			49, 232, 242, 120, 64, 37, 129, 26, 31, 78, 226, 160, 215, 151,
			176, 190, 186, 122, 43, 153, 0, 110, 228, 217, 0, 213, 48, 4,
			124, 39, 32, 102, 130, 197, 135, 204, 183, 9, 244, 165, 28, 138,
			205, 74, 197, 103, 135, 44, 228, 67, 22, 139, 84, 86, 143, 15,
			180, 144, 30, 15, 87, 58, 154, 137, 10, 33, 208, 100, 126, 32,
			100, 28, 116, 70, 50, 224, 17, 208, 200, 135, 145, 96, 16, 68,
			32, 248, 40, 246, 24, 246, 116, 130, 136, 198, 19, 228, 75, 148,
			97, 28, 200, 62, 240, 24, 255, 207, 71, 146, 192, 128, 251, 65,
			55, 240, 168, 162, 80, 6, 26, 51, 24, 178, 120, 16, 72, 201,
			124, 24, 198, 252, 48, 240, 153, 15, 178, 79, 37, 200, 190, 146,
			46, 12, 249, 56, 136, 122, 224, 241, 200, 15, 212, 36, 161, 38,
			17, 24, 48, 185, 73, 8, 168, 159, 47, 95, 99, 76, 0, 239,
			166, 28, 121, 220, 103, 48, 24, 9, 9, 49, 147, 52, 136, 144,
			42, 237, 240, 67, 245, 42, 209, 24, 129, 136, 203, 192, 99, 101,
			144, 253, 64, 64, 24, 8, 169, 40, 204, 175, 24, 249, 175, 177,
			227, 7, 194, 11, 105, 48, 96, 177, 253, 54, 38, 130, 104, 94,
			23, 41, 19, 195, 152, 251, 35, 143, 205, 248, 32, 51, 70, 254,
			175, 248, 32, 144, 72, 231, 115, 111, 52, 96, 145, 164, 233, 38,
			85, 120, 12, 92, 246, 89, 12, 3, 42, 89, 28, 208, 80, 204,
			84, 141, 27, 36, 251, 140, 192, 60, 247, 83, 161, 234, 44, 192,
			153, 138, 112, 68, 7, 76, 49, 52, 111, 91, 17, 159, 189, 67,
			189, 7, 82, 40, 137, 34, 77, 138, 199, 2, 6, 116, 2, 29,
			166, 44, 197, 7, 201, 129, 69, 62, 143, 5, 83, 70, 49, 140,
			249, 128, 75, 6, 90, 39, 82, 128, 207, 226, 224, 144, 249, 208,
			141, 249, 128, 104, 45, 8, 222, 149, 99, 101, 38, 137, 5, 129,
			24, 50, 79, 89, 16, 12, 227, 64, 25, 86, 172, 108, 39, 210,
			86, 36, 4, 242, 78, 160, 125, 223, 109, 65, 171, 177, 211, 126,
			84, 109, 58, 224, 182, 96, 175, 217, 56, 112, 183, 157, 109, 216,
			122, 12, 237, 251, 14, 212, 26, 123, 143, 155, 238, 189, 251, 109,
			184, 223, 216, 221, 118, 154, 45, 168, 214, 183, 161, 214, 168, 183,
			155, 238, 214, 126, 187, 209, 108, 17, 40, 84, 91, 224, 182, 10,
			248, 166, 90, 127, 12, 206, 87, 123, 77, 167, 213, 130, 70, 19,
			220, 135, 123, 187, 174, 179, 13, 143, 170, 205, 102, 181, 222, 118,
			157, 86, 25, 220, 122, 109, 119, 127, 219, 173, 223, 43, 195, 214,
			126, 27, 234, 141, 54, 129, 93, 247, 161, 219, 118, 182, 161, 221,
			40, 227, 178, 111, 206, 131, 198, 14, 60, 116, 154, 181, 251, 213,
			122, 187, 186, 229, 238, 186, 237, 199, 184, 224, 142, 219, 174, 171,
			197, 118, 26, 77, 2, 85, 216, 171, 54, 219, 110, 109, 127, 183,
			218, 132, 189, 253, 230, 94, 163, 229, 128, 146, 108, 219, 109, 213,
			118, 171, 238, 67, 103, 219, 6, 183, 14, 245, 6, 56, 7, 78,
			189, 13, 173, 251, 213, 221, 221, 163, 130, 18, 104, 60, 170, 59,
			77, 197, 253, 188, 152, 176, 229, 192, 174, 91, 221, 218, 117, 212,
			82, 40, 231, 182, 219, 116, 106, 109, 37, 208, 236, 169, 230, 110,
			59, 245, 118, 117, 183, 76, 160, 181, 231, 212, 220, 234, 110, 25,
			156, 175, 156, 135, 123, 187, 213, 230, 227, 114, 66, 180, 229, 252,
			126, 223, 169, 183, 221, 234, 46, 108, 87, 31, 86, 239, 57, 45,
			40, 190, 79, 43, 123, 205, 70, 109, 191, 233, 60, 84, 92, 55,
			118, 160, 181, 191, 213, 106, 187, 237, 253, 182, 3, 247, 26, 141,
			109, 84, 118, 203, 105, 30, 184, 53, 167, 117, 7, 118, 27, 45,
			84, 216, 126, 203, 41, 19, 216, 174, 182, 171, 184, 244, 94, 179,
			177, 227, 182, 91, 119, 212, 243, 214, 126, 203, 69, 197, 185, 245,
			182, 211, 108, 238, 239, 181, 221, 70, 189, 4, 247, 27, 143, 156,
			3, 167, 9, 181, 234, 126, 203, 217, 70, 13, 55, 234, 74, 90,
			101, 43, 78, 163, 249, 88, 145, 85, 122, 192, 29, 40, 195, 163,
			251, 78, 251, 190, 211, 84, 74, 69, 109, 85, 149, 26, 90, 237,
			166, 91, 107, 207, 15, 107, 52, 161, 221, 104, 182, 201, 156, 156,
			80, 119, 238, 237, 186, 247, 156, 122, 205, 81, 175, 27, 138, 204,
			35, 183, 229, 148, 160, 218, 116, 91, 106, 128, 139, 11, 195, 163,
			234, 99, 104, 236, 163, 212, 106, 163, 246, 91, 14, 209, 207, 115,
			166, 91, 198, 253, 4, 119, 7, 170, 219, 7, 174, 226, 60, 25,
			189, 215, 104, 181, 220, 196, 92, 80, 109, 181, 251, 137, 206, 109,
			66, 150, 136, 145, 181, 76, 200, 92, 84, 79, 75, 150, 89, 200,
			220, 33, 39, 72, 118, 105, 89, 63, 234, 206, 207, 50, 14, 118,
			126, 168, 31, 117, 231, 231, 153, 50, 118, 26, 250, 81, 119, 46,
			103, 254, 6, 118, 38, 143, 186, 243, 139, 76, 1, 59, 137, 126,
			212, 157, 87, 51, 87, 176, 243, 115, 253, 168, 59, 139, 153, 79,
			177, 243, 83, 253, 248, 63, 179, 36, 155, 203, 88, 230, 70, 230,
			108, 254, 191, 101, 161, 10, 61, 22, 177, 56, 240, 0, 35, 40,
			12, 152, 16, 180, 199, 116, 8, 152, 240, 17, 120, 52, 130, 152,
			173, 168, 64, 35, 57, 208, 67, 30, 248, 224, 179, 110, 16, 161,
			251, 27, 13, 67, 21, 76, 152, 79, 142, 206, 71, 247, 59, 225,
			163, 24, 170, 123, 174, 176, 161, 10, 114, 50, 12, 60, 26, 2,
			123, 65, 7, 195, 144, 65, 32, 20, 61, 140, 95, 18, 168, 64,
			47, 22, 179, 239, 71, 76, 72, 2, 137, 87, 139, 153, 24, 242,
			72, 173, 60, 25, 162, 235, 163, 145, 162, 167, 130, 79, 159, 251,
			54, 236, 240, 24, 130, 72, 72, 26, 121, 44, 141, 70, 42, 190,
			6, 30, 131, 29, 206, 225, 165, 238, 2, 136, 135, 30, 108, 209,
			184, 248, 26, 214, 176, 17, 106, 148, 84, 108, 26, 197, 145, 128,
			183, 188, 191, 163, 201, 188, 82, 142, 173, 207, 224, 65, 171, 81,
			199, 72, 194, 196, 212, 205, 119, 121, 12, 207, 112, 244, 51, 37,
			153, 214, 5, 14, 228, 157, 239, 152, 39, 225, 217, 203, 87, 207,
			108, 66, 8, 49, 115, 25, 195, 50, 55, 150, 78, 117, 22, 113,
			153, 13, 242, 239, 215, 200, 167, 175, 35, 40, 25, 12, 152, 144,
			116, 48, 124, 27, 138, 186, 67, 78, 180, 211, 49, 214, 69, 242,
			129, 96, 42, 78, 137, 139, 6, 24, 69, 179, 153, 54, 173, 11,
			100, 33, 162, 17, 23, 23, 179, 96, 20, 23, 154, 186, 177, 245,
			183, 141, 227, 161, 215, 233, 41, 201, 20, 126, 173, 255, 72, 248,
			53, 229, 247, 39, 65, 176, 127, 87, 33, 31, 88, 11, 151, 51,
			127, 207, 48, 126, 198, 96, 63, 99, 176, 159, 49, 216, 207, 24,
			236, 103, 12, 246, 51, 6, 251, 255, 136, 193, 166, 208, 72, 61,
			166, 24, 204, 77, 129, 153, 122, 76, 49, 216, 20, 152, 45, 79,
			129, 217, 23, 153, 74, 10, 204, 212, 99, 138, 193, 166, 192, 236,
			234, 20, 152, 21, 103, 192, 76, 61, 254, 231, 95, 33, 6, 91,
			252, 99, 67, 133, 190, 252, 127, 248, 21, 84, 97, 26, 122, 103,
			208, 66, 0, 133, 33, 15, 34, 137, 110, 45, 24, 168, 48, 227,
			179, 33, 139, 124, 22, 73, 13, 135, 38, 186, 255, 7, 30, 161,
			55, 9, 185, 71, 67, 2, 30, 13, 89, 228, 211, 184, 12, 44,
			82, 222, 223, 87, 0, 139, 130, 199, 71, 122, 94, 130, 14, 208,
			151, 118, 99, 234, 205, 34, 70, 250, 66, 5, 4, 5, 21, 176,
			173, 34, 38, 15, 181, 83, 68, 4, 164, 9, 5, 42, 148, 134,
			84, 6, 135, 26, 26, 70, 192, 134, 220, 235, 3, 149, 176, 223,
			174, 193, 32, 240, 35, 244, 232, 60, 34, 240, 128, 70, 35, 21,
			6, 214, 202, 176, 118, 251, 230, 106, 57, 117, 212, 195, 152, 135,
			108, 40, 3, 15, 238, 197, 172, 199, 227, 128, 70, 83, 238, 97,
			220, 15, 188, 62, 176, 23, 146, 41, 158, 208, 65, 31, 51, 170,
			67, 189, 231, 99, 26, 251, 136, 39, 39, 140, 198, 192, 35, 166,
			28, 160, 10, 249, 131, 32, 26, 73, 134, 241, 18, 110, 172, 78,
			229, 11, 121, 212, 179, 97, 151, 209, 225, 76, 228, 152, 65, 65,
			12, 24, 141, 153, 95, 0, 193, 117, 0, 142, 56, 132, 140, 14,
			73, 50, 12, 36, 237, 104, 236, 26, 49, 166, 244, 218, 69, 4,
			42, 89, 60, 84, 177, 85, 7, 244, 145, 80, 81, 137, 194, 215,
			235, 215, 86, 250, 10, 2, 135, 65, 196, 104, 76, 0, 169, 127,
			83, 124, 55, 232, 80, 251, 89, 193, 145, 37, 59, 1, 156, 49,
			162, 156, 64, 96, 76, 128, 213, 213, 213, 181, 21, 252, 215, 94,
			93, 221, 196, 127, 79, 148, 232, 183, 111, 223, 190, 189, 178, 182,
			190, 178, 177, 214, 94, 223, 216, 188, 126, 123, 243, 250, 109, 251,
			118, 250, 243, 196, 134, 173, 9, 81, 27, 41, 227, 192, 147, 138,
			65, 153, 136, 136, 212, 203, 48, 102, 192, 34, 49, 138, 19, 232,
			63, 102, 136, 252, 61, 30, 29, 178, 88, 234, 253, 213, 65, 9,
			190, 110, 238, 212, 8, 108, 108, 108, 220, 158, 201, 50, 30, 143,
			237, 128, 201, 46, 34, 196, 184, 235, 169, 95, 53, 194, 150, 47,
			100, 73, 33, 54, 6, 106, 229, 168, 39, 148, 80, 159, 129, 163,
			179, 0, 65, 72, 250, 8, 107, 155, 80, 227, 131, 225, 72, 178,
			185, 179, 128, 11, 238, 53, 90, 238, 87, 240, 76, 105, 166, 88,
			82, 40, 26, 3, 243, 108, 208, 20, 124, 38, 64, 125, 6, 158,
			5, 147, 79, 147, 13, 46, 226, 244, 250, 254, 238, 110, 169, 116,
			236, 56, 180, 247, 226, 106, 233, 206, 28, 79, 235, 239, 227, 169,
			199, 164, 162, 194, 187, 62, 157, 204, 241, 38, 100, 60, 242, 36,
			46, 112, 72, 67, 144, 135, 201, 138, 71, 134, 127, 33, 15, 203,
			128, 12, 221, 249, 67, 69, 58, 180, 229, 161, 106, 189, 75, 34,
			61, 104, 36, 152, 7, 95, 194, 218, 234, 234, 81, 9, 55, 222,
			42, 225, 163, 32, 218, 88, 135, 103, 247, 152, 108, 77, 132, 100,
			3, 245, 186, 42, 118, 130, 144, 181, 143, 110, 196, 142, 187, 235,
			180, 221, 135, 14, 116, 101, 194, 198, 219, 230, 124, 209, 149, 41,
			167, 251, 110, 189, 125, 227, 26, 200, 192, 123, 46, 224, 46, 20,
			139, 69, 221, 83, 234, 74, 219, 31, 223, 15, 122, 253, 109, 42,
			113, 86, 9, 126, 253, 107, 216, 88, 47, 193, 223, 4, 124, 183,
			203, 199, 233, 171, 84, 111, 149, 10, 84, 21, 191, 62, 31, 11,
			36, 169, 14, 203, 218, 234, 234, 156, 15, 19, 246, 116, 128, 246,
			82, 107, 55, 222, 60, 70, 83, 106, 106, 250, 218, 141, 107, 215,
			174, 221, 220, 184, 177, 58, 115, 27, 29, 214, 229, 49, 131, 253,
			40, 120, 145, 82, 185, 125, 115, 245, 117, 42, 246, 31, 182, 153,
			69, 45, 63, 20, 139, 90, 41, 21, 220, 44, 245, 83, 130, 149,
			121, 118, 222, 99, 193, 138, 142, 82, 87, 74, 103, 121, 142, 14,
			26, 64, 233, 136, 1, 92, 123, 171, 1, 60, 160, 135, 20, 158,
			233, 141, 180, 189, 81, 28, 179, 72, 170, 33, 15, 131, 48, 12,
			196, 156, 1, 40, 111, 10, 3, 236, 133, 187, 240, 246, 9, 239,
			48, 115, 184, 59, 235, 181, 35, 54, 222, 26, 5, 161, 207, 226,
			98, 73, 9, 214, 74, 52, 148, 44, 161, 21, 83, 74, 115, 123,
			0, 53, 166, 174, 101, 15, 34, 169, 36, 79, 70, 106, 209, 19,
			177, 81, 3, 37, 187, 163, 40, 35, 47, 51, 29, 92, 127, 143,
			14, 92, 172, 49, 72, 59, 226, 227, 57, 177, 147, 94, 136, 248,
			24, 238, 194, 145, 49, 239, 148, 116, 198, 248, 251, 69, 142, 248,
			216, 238, 49, 233, 40, 99, 211, 125, 197, 210, 156, 228, 71, 165,
			79, 6, 171, 70, 241, 45, 146, 222, 120, 171, 164, 201, 126, 165,
			56, 3, 246, 38, 178, 175, 19, 137, 35, 134, 54, 191, 81, 197,
			210, 235, 86, 120, 143, 201, 218, 108, 223, 139, 37, 244, 245, 88,
			6, 121, 72, 135, 195, 32, 234, 17, 2, 110, 164, 123, 116, 214,
			94, 70, 24, 48, 167, 167, 201, 16, 67, 221, 17, 224, 162, 67,
			71, 130, 25, 8, 6, 160, 159, 20, 127, 244, 82, 10, 187, 80,
			5, 91, 202, 154, 140, 238, 85, 139, 21, 94, 42, 220, 240, 106,
			229, 229, 128, 71, 178, 255, 106, 229, 165, 79, 39, 175, 218, 47,
			85, 240, 126, 181, 249, 114, 16, 68, 175, 54, 95, 10, 230, 189,
			250, 218, 126, 169, 224, 146, 58, 178, 175, 190, 121, 82, 32, 48,
			238, 179, 152, 129, 158, 173, 8, 209, 112, 76, 39, 2, 216, 11,
			133, 224, 84, 178, 167, 177, 64, 87, 161, 0, 63, 232, 5, 82,
			40, 80, 19, 50, 72, 86, 42, 3, 46, 85, 38, 160, 23, 43,
			3, 174, 86, 198, 96, 139, 75, 34, 46, 249, 129, 197, 124, 101,
			72, 125, 95, 167, 143, 114, 204, 83, 106, 140, 122, 125, 141, 201,
			82, 28, 167, 240, 95, 226, 82, 202, 9, 130, 82, 129, 188, 199,
			97, 52, 68, 152, 144, 78, 45, 6, 54, 179, 147, 206, 181, 227,
			209, 94, 169, 76, 112, 125, 62, 212, 148, 245, 74, 133, 39, 5,
			16, 163, 110, 55, 120, 161, 240, 40, 150, 255, 116, 249, 78, 217,
			1, 34, 209, 98, 97, 191, 93, 43, 148, 238, 28, 233, 37, 26,
			48, 126, 63, 10, 98, 230, 219, 80, 5, 93, 254, 210, 198, 32,
			48, 39, 15, 126, 96, 49, 136, 62, 31, 133, 126, 170, 202, 145,
			96, 136, 38, 139, 84, 76, 87, 243, 161, 51, 33, 138, 141, 146,
			218, 128, 72, 101, 193, 145, 134, 52, 111, 154, 146, 82, 36, 61,
			178, 212, 144, 198, 98, 182, 76, 135, 17, 64, 76, 167, 16, 142,
			231, 177, 161, 132, 14, 151, 125, 92, 83, 205, 213, 69, 131, 84,
			6, 241, 6, 31, 10, 246, 242, 110, 87, 48, 137, 112, 109, 135,
			199, 105, 133, 179, 12, 133, 245, 213, 181, 155, 42, 58, 172, 93,
			111, 175, 174, 109, 110, 172, 110, 174, 93, 183, 87, 215, 158, 20,
			18, 235, 22, 128, 237, 105, 120, 25, 82, 33, 9, 224, 72, 92,
			159, 71, 51, 220, 124, 189, 12, 138, 154, 157, 28, 32, 122, 72,
			91, 94, 28, 12, 101, 89, 161, 221, 35, 80, 141, 130, 10, 143,
			105, 221, 17, 81, 158, 130, 142, 218, 216, 181, 61, 162, 249, 43,
			119, 229, 211, 216, 39, 240, 181, 228, 110, 171, 209, 194, 67, 86,
			44, 29, 3, 80, 237, 1, 255, 33, 8, 67, 138, 167, 139, 69,
			43, 251, 173, 138, 207, 61, 81, 121, 196, 58, 149, 25, 43, 149,
			38, 235, 178, 152, 69, 30, 171, 220, 11, 121, 135, 134, 79, 27,
			200, 131, 168, 40, 134, 42, 115, 139, 148, 200, 180, 132, 235, 166,
			158, 166, 140, 231, 92, 179, 4, 207, 20, 98, 84, 74, 183, 211,
			135, 103, 169, 64, 74, 212, 14, 75, 165, 101, 62, 57, 86, 68,
			2, 95, 63, 19, 50, 238, 226, 212, 57, 137, 184, 39, 236, 161,
			246, 108, 74, 150, 245, 74, 24, 116, 98, 26, 79, 16, 118, 219,
			125, 57, 8, 63, 195, 167, 116, 110, 9, 107, 46, 100, 106, 200,
			233, 34, 98, 200, 60, 184, 186, 252, 120, 101, 121, 176, 178, 236,
			183, 151, 239, 111, 46, 63, 220, 92, 110, 217, 203, 221, 39, 87,
			109, 216, 13, 158, 179, 113, 32, 24, 166, 57, 74, 65, 179, 93,
			26, 9, 166, 169, 61, 224, 62, 69, 99, 189, 42, 224, 235, 103,
			110, 171, 145, 130, 154, 29, 237, 172, 252, 164, 89, 44, 61, 251,
			166, 168, 43, 149, 137, 159, 251, 142, 251, 122, 39, 212, 195, 10,
			230, 11, 116, 24, 224, 134, 164, 189, 58, 139, 208, 188, 86, 222,
			164, 141, 114, 166, 11, 44, 175, 111, 47, 175, 111, 19, 40, 41,
			69, 242, 14, 86, 8, 105, 34, 167, 100, 49, 120, 116, 136, 7,
			132, 119, 245, 85, 1, 213, 71, 45, 61, 102, 66, 187, 229, 169,
			254, 177, 200, 253, 161, 46, 115, 231, 254, 216, 88, 58, 71, 254,
			145, 65, 114, 185, 76, 54, 99, 229, 254, 212, 200, 94, 200, 255,
			153, 1, 205, 89, 134, 155, 218, 62, 239, 162, 201, 163, 142, 69,
			16, 121, 243, 40, 139, 28, 15, 179, 224, 225, 72, 72, 101, 11,
			239, 74, 139, 200, 113, 121, 209, 19, 8, 34, 47, 28, 137, 224,
			80, 37, 138, 167, 200, 130, 98, 111, 1, 249, 251, 32, 109, 26,
			170, 185, 116, 38, 109, 154, 170, 105, 157, 39, 255, 69, 11, 99,
			88, 185, 191, 107, 100, 173, 252, 127, 52, 160, 206, 163, 149, 136,
			245, 116, 30, 124, 36, 155, 166, 105, 214, 168, 18, 201, 227, 179,
			233, 122, 50, 113, 154, 96, 30, 210, 112, 196, 132, 46, 73, 206,
			136, 97, 225, 84, 200, 32, 12, 161, 79, 15, 25, 68, 243, 107,
			34, 233, 100, 34, 209, 217, 155, 78, 208, 187, 60, 86, 137, 113,
			90, 61, 120, 93, 97, 73, 210, 88, 78, 126, 201, 49, 74, 49,
			22, 80, 206, 84, 41, 6, 138, 189, 116, 42, 109, 154, 170, 121,
			246, 220, 244, 38, 227, 95, 93, 33, 43, 65, 212, 141, 105, 133,
			14, 135, 44, 234, 5, 17, 171, 140, 25, 147, 157, 224, 133, 46,
			166, 87, 14, 215, 42, 30, 31, 12, 120, 148, 220, 107, 144, 228,
			181, 125, 184, 150, 127, 223, 37, 72, 97, 172, 239, 60, 154, 42,
			97, 181, 110, 144, 37, 70, 227, 48, 96, 66, 226, 165, 199, 135,
			235, 249, 215, 239, 51, 236, 105, 44, 104, 78, 199, 90, 235, 100,
			49, 84, 17, 75, 226, 149, 200, 187, 103, 37, 35, 11, 55, 200,
			201, 54, 19, 178, 201, 196, 40, 148, 174, 111, 125, 68, 22, 5,
			162, 92, 92, 249, 68, 51, 105, 89, 167, 73, 54, 240, 145, 238,
			137, 102, 54, 240, 11, 223, 147, 15, 14, 104, 28, 208, 72, 90,
			54, 49, 125, 214, 189, 104, 128, 89, 252, 112, 253, 146, 61, 19,
			219, 78, 70, 216, 219, 172, 235, 68, 50, 158, 52, 213, 192, 252,
			13, 178, 148, 118, 88, 103, 137, 249, 156, 77, 146, 181, 212, 163,
			117, 129, 44, 224, 126, 39, 107, 233, 198, 102, 246, 150, 81, 184,
			70, 136, 246, 177, 123, 52, 136, 127, 236, 204, 194, 46, 185, 176,
			53, 234, 181, 99, 234, 61, 15, 162, 158, 66, 136, 60, 98, 145,
			124, 171, 160, 151, 200, 9, 47, 29, 148, 80, 154, 117, 20, 110,
			145, 211, 123, 49, 19, 163, 206, 32, 144, 205, 81, 244, 227, 21,
			246, 229, 51, 114, 234, 128, 197, 126, 224, 201, 150, 164, 114, 36,
			172, 203, 36, 127, 224, 52, 183, 221, 90, 251, 105, 171, 93, 109,
			239, 183, 158, 238, 215, 177, 248, 186, 227, 58, 219, 103, 51, 214,
			105, 66, 246, 235, 206, 87, 123, 78, 173, 237, 108, 159, 37, 214,
			57, 114, 42, 29, 191, 179, 91, 253, 221, 227, 179, 151, 173, 147,
			100, 105, 58, 96, 125, 171, 252, 228, 203, 247, 89, 232, 157, 164,
			99, 216, 121, 240, 159, 62, 33, 139, 86, 46, 151, 97, 6, 249,
			115, 3, 239, 167, 114, 25, 107, 253, 159, 24, 71, 174, 154, 214,
			215, 16, 22, 213, 250, 49, 31, 4, 163, 1, 84, 71, 178, 207,
			99, 97, 191, 229, 206, 105, 95, 160, 47, 77, 42, 251, 179, 27,
			154, 64, 64, 143, 31, 178, 56, 74, 112, 5, 108, 181, 182, 87,
			132, 156, 132, 12, 194, 192, 99, 120, 13, 138, 103, 27, 3, 160,
			130, 175, 163, 200, 79, 235, 104, 187, 110, 205, 169, 183, 28, 232,
			6, 33, 155, 86, 63, 23, 51, 231, 201, 9, 146, 53, 51, 150,
			185, 148, 41, 37, 165, 72, 146, 169, 166, 229, 77, 245, 248, 57,
			86, 34, 115, 167, 50, 231, 141, 252, 69, 168, 38, 181, 38, 197,
			223, 212, 193, 207, 93, 91, 158, 90, 58, 71, 126, 157, 120, 115,
			243, 76, 182, 148, 175, 160, 232, 60, 244, 153, 144, 115, 73, 130,
			228, 218, 153, 248, 44, 101, 16, 233, 218, 132, 156, 212, 238, 116,
			81, 77, 255, 36, 109, 25, 150, 121, 230, 210, 231, 105, 203, 180,
			204, 51, 87, 139, 196, 77, 28, 173, 105, 101, 175, 230, 127, 13,
			110, 66, 143, 71, 225, 100, 62, 250, 160, 78, 20, 72, 213, 149,
			173, 112, 130, 220, 196, 170, 95, 199, 165, 233, 162, 198, 162, 162,
			149, 46, 106, 40, 202, 151, 10, 105, 203, 180, 76, 107, 249, 11,
			242, 23, 6, 201, 46, 100, 172, 220, 197, 204, 23, 70, 254, 95,
			26, 160, 205, 80, 59, 243, 196, 50, 109, 2, 46, 102, 13, 62,
			147, 44, 30, 4, 233, 126, 133, 161, 70, 9, 12, 239, 184, 148,
			167, 16, 122, 159, 25, 28, 234, 153, 26, 214, 179, 23, 92, 71,
			209, 233, 61, 94, 208, 139, 120, 204, 124, 13, 200, 187, 52, 8,
			71, 177, 190, 31, 143, 25, 162, 76, 204, 129, 146, 254, 50, 176,
			67, 22, 65, 208, 133, 0, 153, 72, 169, 49, 191, 164, 247, 105,
			65, 105, 243, 226, 130, 69, 158, 146, 220, 2, 238, 211, 39, 217,
			43, 249, 38, 84, 83, 46, 116, 48, 137, 184, 212, 161, 68, 219,
			33, 138, 105, 19, 104, 171, 86, 32, 180, 150, 241, 186, 10, 17,
			118, 55, 8, 37, 195, 28, 44, 33, 146, 104, 117, 65, 111, 222,
			39, 217, 75, 105, 43, 107, 153, 159, 124, 10, 228, 54, 46, 110,
			88, 230, 229, 172, 149, 47, 235, 147, 112, 172, 78, 112, 235, 70,
			17, 123, 49, 100, 158, 84, 231, 35, 33, 100, 224, 220, 147, 105,
			43, 107, 153, 151, 207, 156, 35, 127, 100, 32, 221, 172, 101, 22,
			178, 191, 200, 11, 52, 190, 148, 80, 159, 10, 13, 221, 83, 90,
			250, 114, 118, 74, 58, 101, 64, 73, 201, 85, 20, 244, 131, 46,
			226, 85, 25, 160, 150, 49, 228, 86, 35, 26, 78, 126, 96, 190,
			114, 247, 137, 99, 214, 38, 96, 163, 59, 153, 178, 167, 68, 43,
			100, 207, 164, 45, 197, 144, 117, 129, 220, 68, 238, 76, 203, 92,
			206, 158, 205, 127, 249, 62, 169, 223, 144, 217, 52, 212, 204, 105,
			43, 107, 153, 203, 167, 206, 144, 34, 201, 230, 12, 43, 87, 202,
			92, 51, 242, 151, 192, 245, 21, 195, 114, 162, 77, 114, 206, 216,
			146, 83, 170, 244, 86, 90, 186, 64, 30, 145, 92, 206, 80, 187,
			95, 206, 94, 200, 63, 64, 69, 29, 177, 76, 237, 128, 109, 2,
			73, 190, 30, 78, 116, 38, 142, 27, 127, 72, 195, 32, 129, 34,
			152, 30, 235, 73, 126, 167, 144, 156, 37, 67, 161, 37, 179, 156,
			93, 74, 91, 134, 101, 150, 79, 156, 73, 91, 166, 101, 150, 173,
			243, 228, 31, 100, 145, 7, 195, 50, 55, 178, 103, 243, 127, 146,
			5, 119, 27, 235, 229, 175, 157, 146, 212, 67, 28, 207, 158, 74,
			168, 142, 188, 9, 34, 208, 113, 120, 123, 171, 156, 92, 4, 39,
			105, 252, 38, 129, 66, 16, 29, 114, 125, 177, 46, 42, 47, 221,
			250, 65, 163, 86, 109, 187, 141, 250, 83, 119, 251, 85, 69, 145,
			17, 149, 151, 251, 205, 221, 167, 78, 171, 86, 221, 115, 182, 159,
			182, 157, 86, 27, 223, 37, 212, 43, 47, 155, 78, 107, 127, 23,
			251, 10, 4, 30, 97, 122, 127, 132, 76, 25, 142, 153, 143, 150,
			54, 157, 137, 155, 155, 224, 56, 252, 82, 70, 37, 41, 115, 108,
			79, 149, 104, 44, 40, 213, 164, 74, 84, 59, 183, 113, 226, 195,
			180, 101, 90, 230, 198, 233, 51, 228, 175, 12, 146, 205, 101, 173,
			220, 102, 230, 55, 70, 254, 47, 13, 72, 140, 242, 232, 37, 209,
			152, 162, 61, 196, 163, 40, 210, 87, 15, 168, 49, 143, 10, 150,
			94, 33, 8, 58, 96, 179, 222, 52, 137, 98, 47, 152, 55, 82,
			182, 31, 68, 179, 211, 160, 168, 137, 50, 238, 84, 250, 173, 14,
			143, 200, 220, 251, 70, 171, 12, 247, 246, 246, 211, 47, 27, 102,
			47, 20, 2, 8, 194, 180, 92, 32, 128, 199, 138, 37, 157, 54,
			133, 180, 151, 6, 18, 101, 17, 155, 75, 103, 200, 223, 81, 80,
			58, 171, 108, 244, 110, 246, 114, 254, 111, 25, 200, 168, 254, 180,
			8, 175, 237, 211, 35, 147, 224, 35, 112, 168, 215, 135, 231, 108,
			178, 162, 13, 115, 72, 131, 248, 136, 26, 136, 74, 237, 233, 64,
			121, 101, 240, 153, 240, 226, 160, 163, 180, 209, 231, 227, 153, 125,
			141, 169, 80, 60, 65, 145, 217, 61, 59, 149, 164, 12, 76, 122,
			118, 41, 217, 151, 44, 70, 167, 187, 217, 95, 164, 45, 195, 50,
			239, 126, 244, 203, 180, 101, 90, 230, 221, 75, 191, 34, 132, 100,
			115, 166, 149, 251, 109, 230, 158, 129, 66, 169, 179, 251, 219, 37,
			139, 252, 142, 228, 114, 166, 146, 169, 150, 61, 151, 255, 13, 52,
			89, 143, 189, 216, 132, 111, 191, 166, 43, 63, 124, 163, 254, 179,
			186, 114, 251, 233, 55, 95, 22, 43, 175, 117, 148, 190, 252, 156,
			192, 67, 250, 2, 66, 22, 245, 100, 127, 19, 110, 92, 75, 216,
			49, 241, 172, 213, 18, 51, 49, 145, 157, 218, 137, 147, 105, 203,
			180, 204, 218, 153, 179, 228, 83, 92, 214, 176, 204, 157, 236, 249,
			188, 117, 132, 210, 250, 245, 27, 83, 82, 202, 226, 118, 166, 164,
			148, 197, 237, 156, 56, 157, 182, 76, 203, 220, 57, 103, 145, 93,
			146, 205, 229, 172, 220, 131, 204, 35, 35, 255, 219, 215, 252, 77,
			103, 212, 3, 153, 160, 68, 152, 2, 62, 192, 140, 241, 200, 187,
			244, 252, 162, 110, 114, 134, 101, 62, 88, 186, 68, 254, 153, 218,
			240, 156, 82, 78, 61, 123, 33, 255, 247, 245, 134, 31, 51, 13,
			60, 30, 235, 79, 191, 252, 233, 69, 149, 10, 135, 169, 249, 150,
			85, 68, 12, 144, 177, 110, 160, 14, 87, 103, 242, 14, 15, 242,
			99, 28, 220, 128, 71, 60, 166, 65, 152, 58, 184, 28, 42, 189,
			158, 104, 42, 135, 74, 175, 39, 14, 46, 135, 74, 175, 91, 231,
			201, 255, 202, 162, 60, 134, 101, 30, 100, 63, 206, 255, 215, 236,
			155, 242, 204, 84, 244, 255, 84, 36, 87, 159, 140, 227, 84, 23,
			8, 72, 133, 73, 190, 161, 9, 116, 113, 110, 198, 10, 77, 238,
			82, 71, 130, 197, 48, 198, 42, 152, 96, 12, 2, 89, 6, 60,
			21, 5, 87, 1, 228, 223, 168, 16, 248, 155, 157, 144, 62, 15,
			34, 38, 68, 65, 127, 109, 55, 79, 27, 25, 32, 51, 14, 134,
			49, 199, 10, 141, 62, 91, 5, 47, 193, 195, 133, 18, 222, 151,
			114, 153, 214, 116, 203, 208, 25, 41, 46, 196, 104, 160, 235, 153,
			10, 205, 38, 159, 180, 48, 127, 238, 102, 88, 81, 187, 42, 224,
			145, 134, 227, 224, 241, 168, 27, 244, 70, 26, 58, 77, 55, 74,
			153, 244, 193, 116, 163, 148, 73, 31, 156, 176, 210, 150, 105, 153,
			7, 191, 248, 136, 252, 158, 100, 115, 11, 86, 238, 73, 134, 25,
			121, 231, 53, 147, 30, 166, 153, 138, 246, 11, 52, 20, 28, 240,
			155, 54, 13, 187, 10, 181, 223, 67, 115, 20, 21, 148, 51, 43,
			212, 14, 240, 57, 65, 90, 185, 5, 195, 50, 159, 44, 125, 68,
			254, 161, 178, 235, 5, 101, 215, 223, 102, 47, 228, 255, 84, 219,
			117, 178, 31, 250, 50, 149, 138, 233, 183, 63, 195, 152, 123, 76,
			136, 68, 198, 185, 181, 127, 164, 169, 134, 35, 47, 88, 241, 14,
			11, 232, 160, 119, 247, 107, 46, 212, 248, 64, 145, 56, 96, 177,
			82, 96, 76, 160, 168, 187, 15, 82, 143, 182, 128, 214, 252, 109,
			162, 164, 5, 180, 230, 111, 19, 107, 94, 64, 107, 254, 214, 58,
			79, 254, 173, 150, 194, 176, 76, 63, 123, 54, 255, 175, 141, 35,
			122, 58, 142, 91, 247, 245, 238, 153, 9, 38, 12, 28, 9, 208,
			105, 206, 147, 138, 178, 73, 0, 160, 240, 82, 13, 125, 186, 215,
			108, 60, 112, 106, 237, 87, 21, 221, 172, 29, 96, 0, 214, 246,
			136, 195, 116, 206, 118, 235, 246, 173, 91, 183, 214, 110, 95, 187,
			177, 113, 235, 250, 181, 149, 181, 149, 238, 237, 107, 55, 55, 214,
			187, 108, 125, 117, 245, 250, 141, 174, 191, 86, 152, 10, 172, 172,
			194, 159, 10, 172, 172, 194, 79, 66, 235, 2, 90, 133, 127, 250,
			204, 180, 106, 241, 71, 231, 201, 173, 183, 229, 132, 120, 183, 31,
			209, 176, 66, 253, 65, 16, 37, 41, 34, 62, 39, 5, 140, 143,
			210, 76, 62, 29, 105, 227, 219, 252, 187, 254, 38, 38, 255, 211,
			138, 36, 133, 191, 48, 200, 47, 157, 23, 67, 30, 203, 57, 88,
			42, 154, 250, 99, 89, 149, 209, 199, 140, 134, 105, 106, 173, 27,
			214, 103, 228, 148, 23, 242, 145, 255, 52, 57, 71, 73, 146, 125,
			18, 59, 247, 116, 159, 117, 145, 124, 224, 83, 73, 5, 147, 23,
			77, 124, 157, 54, 21, 81, 252, 212, 225, 98, 78, 19, 197, 134,
			117, 141, 16, 21, 205, 159, 98, 50, 119, 113, 17, 235, 39, 191,
			152, 175, 101, 76, 203, 51, 205, 19, 50, 125, 92, 255, 142, 44,
			84, 149, 78, 44, 74, 172, 55, 197, 176, 214, 236, 227, 85, 104,
			191, 85, 228, 252, 71, 111, 212, 108, 240, 211, 219, 66, 102, 235,
			198, 147, 107, 63, 101, 43, 239, 224, 243, 176, 243, 224, 191, 159,
			214, 137, 254, 198, 95, 211, 68, 255, 211, 89, 162, 191, 140, 143,
			134, 101, 158, 200, 220, 76, 114, 254, 15, 51, 191, 75, 115, 126,
			245, 120, 157, 100, 23, 51, 86, 238, 116, 230, 162, 145, 47, 1,
			238, 13, 240, 33, 22, 112, 167, 206, 118, 64, 131, 72, 210, 32,
			98, 177, 78, 6, 181, 203, 91, 84, 174, 227, 244, 210, 41, 242,
			47, 76, 146, 91, 196, 236, 242, 227, 236, 65, 254, 31, 155, 240,
			230, 102, 129, 140, 131, 94, 79, 205, 63, 238, 29, 21, 207, 241,
			131, 29, 134, 239, 48, 206, 145, 20, 234, 9, 29, 31, 217, 92,
			80, 64, 251, 75, 46, 63, 244, 217, 192, 176, 41, 202, 208, 249,
			62, 165, 49, 189, 209, 1, 159, 71, 12, 232, 72, 242, 1, 149,
			129, 71, 195, 112, 162, 20, 237, 197, 60, 130, 239, 120, 39, 77,
			115, 155, 123, 181, 163, 169, 174, 10, 69, 212, 123, 174, 116, 27,
			234, 175, 140, 117, 113, 33, 140, 25, 245, 39, 106, 7, 82, 237,
			180, 134, 52, 138, 88, 140, 245, 229, 173, 160, 247, 251, 17, 139,
			39, 54, 184, 18, 124, 206, 68, 116, 85, 194, 152, 199, 207, 85,
			146, 62, 247, 85, 59, 160, 200, 74, 199, 72, 58, 249, 150, 32,
			161, 72, 102, 201, 80, 143, 9, 132, 191, 66, 210, 88, 101, 135,
			202, 207, 138, 145, 215, 159, 209, 137, 3, 148, 124, 204, 240, 43,
			36, 188, 185, 242, 85, 182, 139, 119, 85, 36, 217, 208, 234, 158,
			171, 191, 78, 146, 218, 41, 46, 234, 84, 253, 227, 197, 139, 105,
			43, 107, 153, 31, 255, 114, 61, 109, 153, 150, 249, 241, 221, 38,
			226, 218, 140, 149, 203, 171, 147, 144, 86, 125, 242, 75, 87, 200,
			118, 90, 245, 185, 148, 61, 159, 191, 169, 29, 124, 83, 121, 26,
			27, 212, 198, 206, 182, 46, 189, 13, 64, 55, 148, 230, 185, 60,
			158, 230, 185, 186, 212, 110, 94, 74, 156, 179, 230, 234, 82, 130,
			66, 53, 31, 151, 206, 89, 228, 48, 173, 254, 92, 201, 126, 146,
			15, 166, 74, 78, 62, 192, 58, 106, 56, 243, 118, 163, 19, 30,
			133, 121, 112, 224, 195, 253, 86, 27, 48, 166, 119, 24, 126, 77,
			60, 3, 24, 154, 193, 227, 64, 5, 214, 189, 205, 43, 83, 14,
			85, 248, 184, 114, 226, 163, 185, 82, 209, 149, 95, 230, 201, 135,
			200, 161, 46, 69, 36, 175, 178, 11, 170, 149, 78, 195, 26, 193,
			137, 179, 105, 203, 180, 204, 194, 249, 11, 201, 52, 211, 50, 63,
			203, 158, 79, 94, 153, 11, 170, 149, 78, 83, 71, 247, 179, 169,
			62, 76, 53, 242, 156, 69, 254, 249, 2, 206, 203, 89, 230, 122,
			246, 139, 252, 159, 229, 240, 234, 104, 174, 90, 167, 82, 39, 180,
			88, 174, 19, 187, 169, 202, 193, 73, 42, 223, 152, 164, 238, 98,
			65, 123, 238, 168, 116, 71, 97, 8, 125, 62, 82, 110, 204, 181,
			153, 173, 40, 77, 230, 222, 175, 110, 174, 174, 150, 97, 109, 115,
			117, 21, 108, 219, 38, 208, 80, 54, 54, 14, 208, 71, 177, 9,
			140, 213, 81, 233, 48, 144, 241, 40, 210, 87, 162, 201, 209, 157,
			163, 75, 8, 212, 185, 76, 156, 26, 83, 73, 92, 204, 199, 88,
			174, 161, 32, 152, 202, 217, 100, 114, 185, 150, 126, 138, 134, 183,
			214, 34, 248, 1, 125, 37, 126, 72, 205, 195, 48, 185, 246, 85,
			252, 39, 251, 221, 249, 62, 145, 51, 182, 161, 138, 37, 150, 58,
			63, 68, 63, 93, 158, 173, 163, 166, 211, 32, 18, 176, 134, 236,
			168, 147, 41, 251, 106, 172, 82, 215, 12, 197, 204, 214, 7, 49,
			164, 145, 254, 68, 48, 45, 35, 234, 169, 26, 205, 40, 175, 129,
			82, 139, 62, 141, 253, 196, 214, 213, 60, 149, 227, 171, 163, 56,
			253, 104, 93, 12, 104, 24, 38, 151, 199, 58, 100, 234, 251, 250,
			100, 129, 132, 31, 181, 43, 194, 235, 51, 127, 20, 50, 242, 118,
			87, 137, 233, 184, 154, 156, 108, 118, 74, 156, 71, 76, 216, 100,
			253, 79, 140, 57, 29, 39, 168, 76, 95, 82, 67, 55, 96, 161,
			143, 142, 46, 249, 19, 131, 217, 9, 69, 127, 98, 195, 22, 243,
			40, 254, 125, 79, 63, 16, 100, 38, 160, 238, 58, 66, 42, 230,
			131, 227, 206, 13, 176, 23, 201, 205, 17, 6, 10, 109, 185, 185,
			69, 101, 171, 233, 169, 81, 57, 224, 250, 199, 87, 210, 150, 105,
			153, 235, 159, 47, 167, 16, 236, 255, 4, 0, 0, 255, 255, 107,
			57, 246, 95, 69, 60, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptorpb.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("weetbix.internal.admin.Admin")
	if err != nil {
		panic(err)
	}
	return ret
}