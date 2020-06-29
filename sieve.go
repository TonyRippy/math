package math

// A list of the first 256 prime numbers.
var (
	FirstPrimes = []uint64{
		/* Prime[0] = */ 2,
		/* Prime[1] = */ 3,
		/* Prime[2] = */ 5,
		/* Prime[3] = */ 7,
		/* Prime[4] = */ 11,
		/* Prime[5] = */ 13,
		/* Prime[6] = */ 17,
		/* Prime[7] = */ 19,
		/* Prime[8] = */ 23,
		/* Prime[9] = */ 29,
		/* Prime[10] = */ 31,
		/* Prime[11] = */ 37,
		/* Prime[12] = */ 41,
		/* Prime[13] = */ 43,
		/* Prime[14] = */ 47,
		/* Prime[15] = */ 53,
		/* Prime[16] = */ 59,
		/* Prime[17] = */ 61,
		/* Prime[18] = */ 67,
		/* Prime[19] = */ 71,
		/* Prime[20] = */ 73,
		/* Prime[21] = */ 79,
		/* Prime[22] = */ 83,
		/* Prime[23] = */ 89,
		/* Prime[24] = */ 97,
		/* Prime[25] = */ 101,
		/* Prime[26] = */ 103,
		/* Prime[27] = */ 107,
		/* Prime[28] = */ 109,
		/* Prime[29] = */ 113,
		/* Prime[30] = */ 127,
		/* Prime[31] = */ 131,
		/* Prime[32] = */ 137,
		/* Prime[33] = */ 139,
		/* Prime[34] = */ 149,
		/* Prime[35] = */ 151,
		/* Prime[36] = */ 157,
		/* Prime[37] = */ 163,
		/* Prime[38] = */ 167,
		/* Prime[39] = */ 173,
		/* Prime[40] = */ 179,
		/* Prime[41] = */ 181,
		/* Prime[42] = */ 191,
		/* Prime[43] = */ 193,
		/* Prime[44] = */ 197,
		/* Prime[45] = */ 199,
		/* Prime[46] = */ 211,
		/* Prime[47] = */ 223,
		/* Prime[48] = */ 227,
		/* Prime[49] = */ 229,
		/* Prime[50] = */ 233,
		/* Prime[51] = */ 239,
		/* Prime[52] = */ 241,
		/* Prime[53] = */ 251,
		/* Prime[54] = */ 257,
		/* Prime[55] = */ 263,
		/* Prime[56] = */ 269,
		/* Prime[57] = */ 271,
		/* Prime[58] = */ 277,
		/* Prime[59] = */ 281,
		/* Prime[60] = */ 283,
		/* Prime[61] = */ 293,
		/* Prime[62] = */ 307,
		/* Prime[63] = */ 311,
		/* Prime[64] = */ 313,
		/* Prime[65] = */ 317,
		/* Prime[66] = */ 331,
		/* Prime[67] = */ 337,
		/* Prime[68] = */ 347,
		/* Prime[69] = */ 349,
		/* Prime[70] = */ 353,
		/* Prime[71] = */ 359,
		/* Prime[72] = */ 367,
		/* Prime[73] = */ 373,
		/* Prime[74] = */ 379,
		/* Prime[75] = */ 383,
		/* Prime[76] = */ 389,
		/* Prime[77] = */ 397,
		/* Prime[78] = */ 401,
		/* Prime[79] = */ 409,
		/* Prime[80] = */ 419,
		/* Prime[81] = */ 421,
		/* Prime[82] = */ 431,
		/* Prime[83] = */ 433,
		/* Prime[84] = */ 439,
		/* Prime[85] = */ 443,
		/* Prime[86] = */ 449,
		/* Prime[87] = */ 457,
		/* Prime[88] = */ 461,
		/* Prime[89] = */ 463,
		/* Prime[90] = */ 467,
		/* Prime[91] = */ 479,
		/* Prime[92] = */ 487,
		/* Prime[93] = */ 491,
		/* Prime[94] = */ 499,
		/* Prime[95] = */ 503,
		/* Prime[96] = */ 509,
		/* Prime[97] = */ 521,
		/* Prime[98] = */ 523,
		/* Prime[99] = */ 541,
		/* Prime[100] = */ 547,
		/* Prime[101] = */ 557,
		/* Prime[102] = */ 563,
		/* Prime[103] = */ 569,
		/* Prime[104] = */ 571,
		/* Prime[105] = */ 577,
		/* Prime[106] = */ 587,
		/* Prime[107] = */ 593,
		/* Prime[108] = */ 599,
		/* Prime[109] = */ 601,
		/* Prime[110] = */ 607,
		/* Prime[111] = */ 613,
		/* Prime[112] = */ 617,
		/* Prime[113] = */ 619,
		/* Prime[114] = */ 631,
		/* Prime[115] = */ 641,
		/* Prime[116] = */ 643,
		/* Prime[117] = */ 647,
		/* Prime[118] = */ 653,
		/* Prime[119] = */ 659,
		/* Prime[120] = */ 661,
		/* Prime[121] = */ 673,
		/* Prime[122] = */ 677,
		/* Prime[123] = */ 683,
		/* Prime[124] = */ 691,
		/* Prime[125] = */ 701,
		/* Prime[126] = */ 709,
		/* Prime[127] = */ 719,
		/* Prime[128] = */ 727,
		/* Prime[129] = */ 733,
		/* Prime[130] = */ 739,
		/* Prime[131] = */ 743,
		/* Prime[132] = */ 751,
		/* Prime[133] = */ 757,
		/* Prime[134] = */ 761,
		/* Prime[135] = */ 769,
		/* Prime[136] = */ 773,
		/* Prime[137] = */ 787,
		/* Prime[138] = */ 797,
		/* Prime[139] = */ 809,
		/* Prime[140] = */ 811,
		/* Prime[141] = */ 821,
		/* Prime[142] = */ 823,
		/* Prime[143] = */ 827,
		/* Prime[144] = */ 829,
		/* Prime[145] = */ 839,
		/* Prime[146] = */ 853,
		/* Prime[147] = */ 857,
		/* Prime[148] = */ 859,
		/* Prime[149] = */ 863,
		/* Prime[150] = */ 877,
		/* Prime[151] = */ 881,
		/* Prime[152] = */ 883,
		/* Prime[153] = */ 887,
		/* Prime[154] = */ 907,
		/* Prime[155] = */ 911,
		/* Prime[156] = */ 919,
		/* Prime[157] = */ 929,
		/* Prime[158] = */ 937,
		/* Prime[159] = */ 941,
		/* Prime[160] = */ 947,
		/* Prime[161] = */ 953,
		/* Prime[162] = */ 967,
		/* Prime[163] = */ 971,
		/* Prime[164] = */ 977,
		/* Prime[165] = */ 983,
		/* Prime[166] = */ 991,
		/* Prime[167] = */ 997,
		/* Prime[168] = */ 1009,
		/* Prime[169] = */ 1013,
		/* Prime[170] = */ 1019,
		/* Prime[171] = */ 1021,
		/* Prime[172] = */ 1031,
		/* Prime[173] = */ 1033,
		/* Prime[174] = */ 1039,
		/* Prime[175] = */ 1049,
		/* Prime[176] = */ 1051,
		/* Prime[177] = */ 1061,
		/* Prime[178] = */ 1063,
		/* Prime[179] = */ 1069,
		/* Prime[180] = */ 1087,
		/* Prime[181] = */ 1091,
		/* Prime[182] = */ 1093,
		/* Prime[183] = */ 1097,
		/* Prime[184] = */ 1103,
		/* Prime[185] = */ 1109,
		/* Prime[186] = */ 1117,
		/* Prime[187] = */ 1123,
		/* Prime[188] = */ 1129,
		/* Prime[189] = */ 1151,
		/* Prime[190] = */ 1153,
		/* Prime[191] = */ 1163,
		/* Prime[192] = */ 1171,
		/* Prime[193] = */ 1181,
		/* Prime[194] = */ 1187,
		/* Prime[195] = */ 1193,
		/* Prime[196] = */ 1201,
		/* Prime[197] = */ 1213,
		/* Prime[198] = */ 1217,
		/* Prime[199] = */ 1223,
		/* Prime[200] = */ 1229,
		/* Prime[201] = */ 1231,
		/* Prime[202] = */ 1237,
		/* Prime[203] = */ 1249,
		/* Prime[204] = */ 1259,
		/* Prime[205] = */ 1277,
		/* Prime[206] = */ 1279,
		/* Prime[207] = */ 1283,
		/* Prime[208] = */ 1289,
		/* Prime[209] = */ 1291,
		/* Prime[210] = */ 1297,
		/* Prime[211] = */ 1301,
		/* Prime[212] = */ 1303,
		/* Prime[213] = */ 1307,
		/* Prime[214] = */ 1319,
		/* Prime[215] = */ 1321,
		/* Prime[216] = */ 1327,
		/* Prime[217] = */ 1361,
		/* Prime[218] = */ 1367,
		/* Prime[219] = */ 1373,
		/* Prime[220] = */ 1381,
		/* Prime[221] = */ 1399,
		/* Prime[222] = */ 1409,
		/* Prime[223] = */ 1423,
		/* Prime[224] = */ 1427,
		/* Prime[225] = */ 1429,
		/* Prime[226] = */ 1433,
		/* Prime[227] = */ 1439,
		/* Prime[228] = */ 1447,
		/* Prime[229] = */ 1451,
		/* Prime[230] = */ 1453,
		/* Prime[231] = */ 1459,
		/* Prime[232] = */ 1471,
		/* Prime[233] = */ 1481,
		/* Prime[234] = */ 1483,
		/* Prime[235] = */ 1487,
		/* Prime[236] = */ 1489,
		/* Prime[237] = */ 1493,
		/* Prime[238] = */ 1499,
		/* Prime[239] = */ 1511,
		/* Prime[240] = */ 1523,
		/* Prime[241] = */ 1531,
		/* Prime[242] = */ 1543,
		/* Prime[243] = */ 1549,
		/* Prime[244] = */ 1553,
		/* Prime[245] = */ 1559,
		/* Prime[246] = */ 1567,
		/* Prime[247] = */ 1571,
		/* Prime[248] = */ 1579,
		/* Prime[249] = */ 1583,
		/* Prime[250] = */ 1597,
		/* Prime[251] = */ 1601,
		/* Prime[252] = */ 1607,
		/* Prime[253] = */ 1609,
		/* Prime[254] = */ 1613,
		/* Prime[255] = */ 1619,
	}

	firstMarks = []uint64{  
		/* Prime[0] = 2, FirstMark = */ 1620,
		/* Prime[1] = 3, FirstMark = */ 1620,
		/* Prime[2] = 5, FirstMark = */ 1620,
		/* Prime[3] = 7, FirstMark = */ 1624,
		/* Prime[4] = 11, FirstMark = */ 1628,
		/* Prime[5] = 13, FirstMark = */ 1625,
		/* Prime[6] = 17, FirstMark = */ 1632,
		/* Prime[7] = 19, FirstMark = */ 1634,
		/* Prime[8] = 23, FirstMark = */ 1633,
		/* Prime[9] = 29, FirstMark = */ 1624,
		/* Prime[10] = 31, FirstMark = */ 1643,
		/* Prime[11] = 37, FirstMark = */ 1628,
		/* Prime[12] = 41, FirstMark = */ 1681,
		/* Prime[13] = 43, FirstMark = */ 1849,
		/* Prime[14] = 47, FirstMark = */ 2209,
		/* Prime[15] = 53, FirstMark = */ 2809,
		/* Prime[16] = 59, FirstMark = */ 3481,
		/* Prime[17] = 61, FirstMark = */ 3721,
		/* Prime[18] = 67, FirstMark = */ 4489,
		/* Prime[19] = 71, FirstMark = */ 5041,
		/* Prime[20] = 73, FirstMark = */ 5329,
		/* Prime[21] = 79, FirstMark = */ 6241,
		/* Prime[22] = 83, FirstMark = */ 6889,
		/* Prime[23] = 89, FirstMark = */ 7921,
		/* Prime[24] = 97, FirstMark = */ 9409,
		/* Prime[25] = 101, FirstMark = */ 10201,
		/* Prime[26] = 103, FirstMark = */ 10609,
		/* Prime[27] = 107, FirstMark = */ 11449,
		/* Prime[28] = 109, FirstMark = */ 11881,
		/* Prime[29] = 113, FirstMark = */ 12769,
		/* Prime[30] = 127, FirstMark = */ 16129,
		/* Prime[31] = 131, FirstMark = */ 17161,
		/* Prime[32] = 137, FirstMark = */ 18769,
		/* Prime[33] = 139, FirstMark = */ 19321,
		/* Prime[34] = 149, FirstMark = */ 22201,
		/* Prime[35] = 151, FirstMark = */ 22801,
		/* Prime[36] = 157, FirstMark = */ 24649,
		/* Prime[37] = 163, FirstMark = */ 26569,
		/* Prime[38] = 167, FirstMark = */ 27889,
		/* Prime[39] = 173, FirstMark = */ 29929,
		/* Prime[40] = 179, FirstMark = */ 32041,
		/* Prime[41] = 181, FirstMark = */ 32761,
		/* Prime[42] = 191, FirstMark = */ 36481,
		/* Prime[43] = 193, FirstMark = */ 37249,
		/* Prime[44] = 197, FirstMark = */ 38809,
		/* Prime[45] = 199, FirstMark = */ 39601,
		/* Prime[46] = 211, FirstMark = */ 44521,
		/* Prime[47] = 223, FirstMark = */ 49729,
		/* Prime[48] = 227, FirstMark = */ 51529,
		/* Prime[49] = 229, FirstMark = */ 52441,
		/* Prime[50] = 233, FirstMark = */ 54289,
		/* Prime[51] = 239, FirstMark = */ 57121,
		/* Prime[52] = 241, FirstMark = */ 58081,
		/* Prime[53] = 251, FirstMark = */ 63001,
		/* Prime[54] = 257, FirstMark = */ 66049,
		/* Prime[55] = 263, FirstMark = */ 69169,
		/* Prime[56] = 269, FirstMark = */ 72361,
		/* Prime[57] = 271, FirstMark = */ 73441,
		/* Prime[58] = 277, FirstMark = */ 76729,
		/* Prime[59] = 281, FirstMark = */ 78961,
		/* Prime[60] = 283, FirstMark = */ 80089,
		/* Prime[61] = 293, FirstMark = */ 85849,
		/* Prime[62] = 307, FirstMark = */ 94249,
		/* Prime[63] = 311, FirstMark = */ 96721,
		/* Prime[64] = 313, FirstMark = */ 97969,
		/* Prime[65] = 317, FirstMark = */ 100489,
		/* Prime[66] = 331, FirstMark = */ 109561,
		/* Prime[67] = 337, FirstMark = */ 113569,
		/* Prime[68] = 347, FirstMark = */ 120409,
		/* Prime[69] = 349, FirstMark = */ 121801,
		/* Prime[70] = 353, FirstMark = */ 124609,
		/* Prime[71] = 359, FirstMark = */ 128881,
		/* Prime[72] = 367, FirstMark = */ 134689,
		/* Prime[73] = 373, FirstMark = */ 139129,
		/* Prime[74] = 379, FirstMark = */ 143641,
		/* Prime[75] = 383, FirstMark = */ 146689,
		/* Prime[76] = 389, FirstMark = */ 151321,
		/* Prime[77] = 397, FirstMark = */ 157609,
		/* Prime[78] = 401, FirstMark = */ 160801,
		/* Prime[79] = 409, FirstMark = */ 167281,
		/* Prime[80] = 419, FirstMark = */ 175561,
		/* Prime[81] = 421, FirstMark = */ 177241,
		/* Prime[82] = 431, FirstMark = */ 185761,
		/* Prime[83] = 433, FirstMark = */ 187489,
		/* Prime[84] = 439, FirstMark = */ 192721,
		/* Prime[85] = 443, FirstMark = */ 196249,
		/* Prime[86] = 449, FirstMark = */ 201601,
		/* Prime[87] = 457, FirstMark = */ 208849,
		/* Prime[88] = 461, FirstMark = */ 212521,
		/* Prime[89] = 463, FirstMark = */ 214369,
		/* Prime[90] = 467, FirstMark = */ 218089,
		/* Prime[91] = 479, FirstMark = */ 229441,
		/* Prime[92] = 487, FirstMark = */ 237169,
		/* Prime[93] = 491, FirstMark = */ 241081,
		/* Prime[94] = 499, FirstMark = */ 249001,
		/* Prime[95] = 503, FirstMark = */ 253009,
		/* Prime[96] = 509, FirstMark = */ 259081,
		/* Prime[97] = 521, FirstMark = */ 271441,
		/* Prime[98] = 523, FirstMark = */ 273529,
		/* Prime[99] = 541, FirstMark = */ 292681,
		/* Prime[100] = 547, FirstMark = */ 299209,
		/* Prime[101] = 557, FirstMark = */ 310249,
		/* Prime[102] = 563, FirstMark = */ 316969,
		/* Prime[103] = 569, FirstMark = */ 323761,
		/* Prime[104] = 571, FirstMark = */ 326041,
		/* Prime[105] = 577, FirstMark = */ 332929,
		/* Prime[106] = 587, FirstMark = */ 344569,
		/* Prime[107] = 593, FirstMark = */ 351649,
		/* Prime[108] = 599, FirstMark = */ 358801,
		/* Prime[109] = 601, FirstMark = */ 361201,
		/* Prime[110] = 607, FirstMark = */ 368449,
		/* Prime[111] = 613, FirstMark = */ 375769,
		/* Prime[112] = 617, FirstMark = */ 380689,
		/* Prime[113] = 619, FirstMark = */ 383161,
		/* Prime[114] = 631, FirstMark = */ 398161,
		/* Prime[115] = 641, FirstMark = */ 410881,
		/* Prime[116] = 643, FirstMark = */ 413449,
		/* Prime[117] = 647, FirstMark = */ 418609,
		/* Prime[118] = 653, FirstMark = */ 426409,
		/* Prime[119] = 659, FirstMark = */ 434281,
		/* Prime[120] = 661, FirstMark = */ 436921,
		/* Prime[121] = 673, FirstMark = */ 452929,
		/* Prime[122] = 677, FirstMark = */ 458329,
		/* Prime[123] = 683, FirstMark = */ 466489,
		/* Prime[124] = 691, FirstMark = */ 477481,
		/* Prime[125] = 701, FirstMark = */ 491401,
		/* Prime[126] = 709, FirstMark = */ 502681,
		/* Prime[127] = 719, FirstMark = */ 516961,
		/* Prime[128] = 727, FirstMark = */ 528529,
		/* Prime[129] = 733, FirstMark = */ 537289,
		/* Prime[130] = 739, FirstMark = */ 546121,
		/* Prime[131] = 743, FirstMark = */ 552049,
		/* Prime[132] = 751, FirstMark = */ 564001,
		/* Prime[133] = 757, FirstMark = */ 573049,
		/* Prime[134] = 761, FirstMark = */ 579121,
		/* Prime[135] = 769, FirstMark = */ 591361,
		/* Prime[136] = 773, FirstMark = */ 597529,
		/* Prime[137] = 787, FirstMark = */ 619369,
		/* Prime[138] = 797, FirstMark = */ 635209,
		/* Prime[139] = 809, FirstMark = */ 654481,
		/* Prime[140] = 811, FirstMark = */ 657721,
		/* Prime[141] = 821, FirstMark = */ 674041,
		/* Prime[142] = 823, FirstMark = */ 677329,
		/* Prime[143] = 827, FirstMark = */ 683929,
		/* Prime[144] = 829, FirstMark = */ 687241,
		/* Prime[145] = 839, FirstMark = */ 703921,
		/* Prime[146] = 853, FirstMark = */ 727609,
		/* Prime[147] = 857, FirstMark = */ 734449,
		/* Prime[148] = 859, FirstMark = */ 737881,
		/* Prime[149] = 863, FirstMark = */ 744769,
		/* Prime[150] = 877, FirstMark = */ 769129,
		/* Prime[151] = 881, FirstMark = */ 776161,
		/* Prime[152] = 883, FirstMark = */ 779689,
		/* Prime[153] = 887, FirstMark = */ 786769,
		/* Prime[154] = 907, FirstMark = */ 822649,
		/* Prime[155] = 911, FirstMark = */ 829921,
		/* Prime[156] = 919, FirstMark = */ 844561,
		/* Prime[157] = 929, FirstMark = */ 863041,
		/* Prime[158] = 937, FirstMark = */ 877969,
		/* Prime[159] = 941, FirstMark = */ 885481,
		/* Prime[160] = 947, FirstMark = */ 896809,
		/* Prime[161] = 953, FirstMark = */ 908209,
		/* Prime[162] = 967, FirstMark = */ 935089,
		/* Prime[163] = 971, FirstMark = */ 942841,
		/* Prime[164] = 977, FirstMark = */ 954529,
		/* Prime[165] = 983, FirstMark = */ 966289,
		/* Prime[166] = 991, FirstMark = */ 982081,
		/* Prime[167] = 997, FirstMark = */ 994009,
		/* Prime[168] = 1009, FirstMark = */ 1018081,
		/* Prime[169] = 1013, FirstMark = */ 1026169,
		/* Prime[170] = 1019, FirstMark = */ 1038361,
		/* Prime[171] = 1021, FirstMark = */ 1042441,
		/* Prime[172] = 1031, FirstMark = */ 1062961,
		/* Prime[173] = 1033, FirstMark = */ 1067089,
		/* Prime[174] = 1039, FirstMark = */ 1079521,
		/* Prime[175] = 1049, FirstMark = */ 1100401,
		/* Prime[176] = 1051, FirstMark = */ 1104601,
		/* Prime[177] = 1061, FirstMark = */ 1125721,
		/* Prime[178] = 1063, FirstMark = */ 1129969,
		/* Prime[179] = 1069, FirstMark = */ 1142761,
		/* Prime[180] = 1087, FirstMark = */ 1181569,
		/* Prime[181] = 1091, FirstMark = */ 1190281,
		/* Prime[182] = 1093, FirstMark = */ 1194649,
		/* Prime[183] = 1097, FirstMark = */ 1203409,
		/* Prime[184] = 1103, FirstMark = */ 1216609,
		/* Prime[185] = 1109, FirstMark = */ 1229881,
		/* Prime[186] = 1117, FirstMark = */ 1247689,
		/* Prime[187] = 1123, FirstMark = */ 1261129,
		/* Prime[188] = 1129, FirstMark = */ 1274641,
		/* Prime[189] = 1151, FirstMark = */ 1324801,
		/* Prime[190] = 1153, FirstMark = */ 1329409,
		/* Prime[191] = 1163, FirstMark = */ 1352569,
		/* Prime[192] = 1171, FirstMark = */ 1371241,
		/* Prime[193] = 1181, FirstMark = */ 1394761,
		/* Prime[194] = 1187, FirstMark = */ 1408969,
		/* Prime[195] = 1193, FirstMark = */ 1423249,
		/* Prime[196] = 1201, FirstMark = */ 1442401,
		/* Prime[197] = 1213, FirstMark = */ 1471369,
		/* Prime[198] = 1217, FirstMark = */ 1481089,
		/* Prime[199] = 1223, FirstMark = */ 1495729,
		/* Prime[200] = 1229, FirstMark = */ 1510441,
		/* Prime[201] = 1231, FirstMark = */ 1515361,
		/* Prime[202] = 1237, FirstMark = */ 1530169,
		/* Prime[203] = 1249, FirstMark = */ 1560001,
		/* Prime[204] = 1259, FirstMark = */ 1585081,
		/* Prime[205] = 1277, FirstMark = */ 1630729,
		/* Prime[206] = 1279, FirstMark = */ 1635841,
		/* Prime[207] = 1283, FirstMark = */ 1646089,
		/* Prime[208] = 1289, FirstMark = */ 1661521,
		/* Prime[209] = 1291, FirstMark = */ 1666681,
		/* Prime[210] = 1297, FirstMark = */ 1682209,
		/* Prime[211] = 1301, FirstMark = */ 1692601,
		/* Prime[212] = 1303, FirstMark = */ 1697809,
		/* Prime[213] = 1307, FirstMark = */ 1708249,
		/* Prime[214] = 1319, FirstMark = */ 1739761,
		/* Prime[215] = 1321, FirstMark = */ 1745041,
		/* Prime[216] = 1327, FirstMark = */ 1760929,
		/* Prime[217] = 1361, FirstMark = */ 1852321,
		/* Prime[218] = 1367, FirstMark = */ 1868689,
		/* Prime[219] = 1373, FirstMark = */ 1885129,
		/* Prime[220] = 1381, FirstMark = */ 1907161,
		/* Prime[221] = 1399, FirstMark = */ 1957201,
		/* Prime[222] = 1409, FirstMark = */ 1985281,
		/* Prime[223] = 1423, FirstMark = */ 2024929,
		/* Prime[224] = 1427, FirstMark = */ 2036329,
		/* Prime[225] = 1429, FirstMark = */ 2042041,
		/* Prime[226] = 1433, FirstMark = */ 2053489,
		/* Prime[227] = 1439, FirstMark = */ 2070721,
		/* Prime[228] = 1447, FirstMark = */ 2093809,
		/* Prime[229] = 1451, FirstMark = */ 2105401,
		/* Prime[230] = 1453, FirstMark = */ 2111209,
		/* Prime[231] = 1459, FirstMark = */ 2128681,
		/* Prime[232] = 1471, FirstMark = */ 2163841,
		/* Prime[233] = 1481, FirstMark = */ 2193361,
		/* Prime[234] = 1483, FirstMark = */ 2199289,
		/* Prime[235] = 1487, FirstMark = */ 2211169,
		/* Prime[236] = 1489, FirstMark = */ 2217121,
		/* Prime[237] = 1493, FirstMark = */ 2229049,
		/* Prime[238] = 1499, FirstMark = */ 2247001,
		/* Prime[239] = 1511, FirstMark = */ 2283121,
		/* Prime[240] = 1523, FirstMark = */ 2319529,
		/* Prime[241] = 1531, FirstMark = */ 2343961,
		/* Prime[242] = 1543, FirstMark = */ 2380849,
		/* Prime[243] = 1549, FirstMark = */ 2399401,
		/* Prime[244] = 1553, FirstMark = */ 2411809,
		/* Prime[245] = 1559, FirstMark = */ 2430481,
		/* Prime[246] = 1567, FirstMark = */ 2455489,
		/* Prime[247] = 1571, FirstMark = */ 2468041,
		/* Prime[248] = 1579, FirstMark = */ 2493241,
		/* Prime[249] = 1583, FirstMark = */ 2505889,
		/* Prime[250] = 1597, FirstMark = */ 2550409,
		/* Prime[251] = 1601, FirstMark = */ 2563201,
		/* Prime[252] = 1607, FirstMark = */ 2582449,
		/* Prime[253] = 1609, FirstMark = */ 2588881,
		/* Prime[254] = 1613, FirstMark = */ 2601769,
		/* Prime[255] = 1619, FirstMark = */ 2621161,
	}
)

// An implementation of the Sieve of Eratosthenes.
// It is a stateful object that helps generate prime numbers in sequence.
type Sieve interface {
	// Next generates the next prime number in the sequence.
	Next() uint64
}

type sieve struct {
	n uint64
	primes []uint64
	marks []uint64
}

// NewSieve creates a new sieve that starts generating primes at 2.
func NewSieve() Sieve {
	return &sieve{0, nil, nil}
}

func (s *sieve) Next() uint64 {
	if s.primes == nil {
		p := FirstPrimes[s.n]
		s.n += 1		
		if s.n == 256 {
			// We know we can skip even numbers, so we increment n by 2
			// and leave out the first prime & mark entries.
			s.n = p + 2
			s.primes = FirstPrimes[1:]
			s.marks = firstMarks[1:]
		}
		return p
	}
	// The sieve stage. Keep incrementing 'n' until we find a value
	// that passes all the incremental filters.
	n := s.n
	for i := 0; i < len(s.primes); {
		for ; s.marks[i] < n;  {
			s.marks[i] += s.primes[i]
		}
		if s.marks[i] == n {
			// This value of n was filtered out. Increment n and try again.
			n += 2
			i = 0
			continue
		}
		// n passed this filter, try the next.
		i += 1
	}
	// n is prime! Add it to the lists before returning.
	s.n = n + 2
	s.primes = append(s.primes, n)
	s.marks = append(s.marks, n * n)
	return n
}
