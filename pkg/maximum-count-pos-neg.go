package pkg

import "fmt"

func findfirstByCondition(isGood func(int) bool, xs []int) int {
	low, high := 0, len(xs)-1
	for {
		mid := (low + high) / 2
		// fmt.Printf("iter: %d, %d, %d, %d, %t\n", low, mid, high, xs[mid], isGood(xs[mid]))
		if high <= low {
			if isGood(xs[mid]) {
				return mid
			}
			return -1
		} else if isGood(xs[mid]) {
			high = mid
		} else {
			low = mid + 1
		}
	}
}

func isPos(x int) bool {
	return x > 0
}

func isNonNeg(x int) bool {
	return x >= 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumCount(nums []int) int {
	pos, zeroes, neg := 0, 0, 0
	firstPos := findfirstByCondition(isPos, nums)
	if firstPos >= 0 {
		pos = len(nums) - firstPos
	}
	firstNonNeg := findfirstByCondition(isNonNeg, nums)
	if firstNonNeg >= 0 && nums[firstNonNeg] == 0 {
		zeroes = len(nums) - (firstNonNeg + pos)
	}
	neg = len(nums) - pos - zeroes
	fmt.Printf("%d, %d; %d, %d, %d\n", firstPos, firstNonNeg, pos, zeroes, neg)
	return max(pos, neg)
}

func MaximumCountMain() {
	cases := [][]int{
		{-2, -1, -1, 1, 2, 3},
		{-3, -2, -1, 0, 0, 1, 2},
		{5, 20, 66, 1314},
		{-1, 0, 1},
		{-1, -1, 0, 1},
		{-1, 0, 1, 1},
		{-1, 0, 1},
		{-1, 0, 1},
		{-1993, -1991, -1985, -1983, -1978, -1956, -1943, -1935, -1929, -1928, -1925, -1924, -1896, -1894, -1893, -1884, -1882, -1871, -1861, -1860, -1858, -1849, -1841, -1837, -1829, -1828, -1811, -1808, -1804, -1804, -1800, -1800, -1796, -1791, -1791, -1777, -1764, -1755, -1745, -1726, -1717, -1708, -1697, -1694, -1694, -1691, -1679, -1677, -1677, -1660, -1654, -1642, -1623, -1621, -1593, -1585, -1581, -1580, -1573, -1558, -1549, -1537, -1527, -1523, -1515, -1513, -1508, -1505, -1497, -1478, -1468, -1463, -1447, -1439, -1429, -1425, -1420, -1412, -1411, -1398, -1396, -1393, -1393, -1387, -1382, -1381, -1378, -1372, -1371, -1341, -1324, -1310, -1295, -1286, -1280, -1279, -1277, -1273, -1257, -1252, -1249, -1244, -1212, -1201, -1197, -1185, -1166, -1140, -1137, -1124, -1118, -1100, -1093, -1084, -1078, -1075, -1064, -1054, -1046, -1039, -1010, -1001, -1000, -987, -949, -946, -942, -927, -923, -921, -921, -910, -906, -897, -895, -894, -893, -887, -874, -868, -861, -852, -849, -848, -848, -846, -840, -828, -825, -822, -821, -819, -807, -801, -800, -788, -785, -785, -784, -777, -764, -757, -754, -734, -730, -712, -707, -706, -703, -682, -681, -680, -676, -668, -667, -662, -654, -647, -640, -624, -607, -601, -598, -597, -586, -583, -582, -574, -572, -556, -547, -539, -536, -517, -500, -496, -487, -478, -470, -469, -461, -449, -449, -435, -427, -420, -416, -403, -388, -386, -386, -383, -376, -367, -365, -361, -340, -338, -337, -335, -334, -325, -322, -303, -295, -283, -273, -271, -263, -263, -256, -251, -243, -229, -223, -212, -207, -201, -199, -199, -198, -186, -175, -162, -161, -161, -146, -140, -138, -135, -127, -125, -113, -96, -90, -85, -77, -75, -60, -52, -26, -11, -7, 0, 5, 10, 29, 46, 67, 67, 68, 70, 92, 106, 110, 138, 147, 149, 155, 159, 178, 178, 191, 196, 197, 200, 210, 211, 226, 233, 239, 248, 258, 265, 267, 271, 285, 288, 293, 311, 318, 320, 324, 325, 327, 333, 340, 374, 375, 379, 386, 390, 393, 403, 410, 415, 418, 419, 462, 482, 486, 486, 489, 493, 501, 508, 524, 527, 532, 532, 534, 557, 568, 573, 580, 598, 603, 604, 608, 609, 620, 625, 626, 630, 663, 664, 698, 706, 716, 727, 739, 751, 752, 759, 761, 781, 785, 807, 807, 816, 820, 820, 841, 844, 849, 856, 859, 877, 877, 882, 889, 890, 917, 918, 920, 923, 926, 929, 937, 950, 956, 956, 971, 974, 978, 989, 998, 999, 1005, 1016, 1022, 1056, 1058, 1059, 1068, 1082, 1085, 1085, 1091, 1097, 1104, 1117, 1124, 1126, 1130, 1134, 1136, 1143, 1152, 1173, 1192, 1197, 1203, 1203, 1216, 1217, 1235, 1248, 1268, 1272, 1275, 1277, 1280, 1286, 1305, 1314, 1340, 1356, 1361, 1363, 1368, 1394, 1398, 1402, 1404, 1423, 1428, 1431, 1435, 1456, 1457, 1468, 1475, 1484, 1498, 1498, 1534, 1542, 1553, 1570, 1571, 1590, 1594, 1599, 1611, 1625, 1633, 1649, 1666, 1670, 1684, 1687, 1688, 1701, 1706, 1709, 1712, 1712, 1726, 1732, 1744, 1753, 1755, 1775, 1780, 1786, 1791, 1804, 1823, 1834, 1847, 1852, 1865, 1865, 1878, 1883, 1890, 1896, 1908, 1915, 1927, 1932, 1936, 1948, 1950, 1968, 1968, 1990, 1991, 1991},
		{-1993, -1991, -1985, -1983, -1978, -1956, -1943, -1935, -1929, -1928, -1925, -1924, -1896, -1894, -1893, -1884, -1882, -1871, -1861, -1860, -1858, -1849, -1841, -1837, -1829, -1828, -1811, -1808, -1804, -1804, -1800, -1800, -1796, -1791, -1791, -1777, -1764, -1755, -1745, -1726, -1717, -1708, -1697, -1694, -1694, -1691, -1679, -1677, -1677, -1660, -1654, -1642, -1623, -1621, -1593, -1585, -1581, -1580, -1573, -1558, -1549, -1537, -1527, -1523, -1515, -1513, -1508, -1505, -1497, -1478, -1468, -1463, -1447, -1439, -1429, -1425, -1420, -1412, -1411, -1398, -1396, -1393, -1393, -1387, -1382, -1381, -1378, -1372, -1371, -1341, -1324, -1310, -1295, -1286, -1280, -1279, -1277, -1273, -1257, -1252, -1249, -1244, -1212, -1201, -1197, -1185, -1166, -1140, -1137, -1124, -1118, -1100, -1093, -1084, -1078, -1075, -1064, -1054, -1046, -1039, -1010, -1001, -1000, -987, -949, -946, -942, -927, -923, -921, -921, -910, -906, -897, -895, -894, -893, -887, -874, -868, -861, -852, -849, -848, -848, -846, -840, -828, -825, -822, -821, -819, -807, -801, -800, -788, -785, -785, -784, -777, -764, -757, -754, -734, -730, -712, -707, -706, -703, -682, -681, -680, -676, -668, -667, -662, -654, -647, -640, -624, -607, -601, -598, -597, -586, -583, -582, -574, -572, -556, -547, -539, -536, -517, -500, -496, -487, -478, -470, -469, -461, -449, -449, -435, -427, -420, -416, -403, -388, -386, -386, -383, -376, -367, -365, -361, -340, -338, -337, -335, -334, -325, -322, -303, -295, -283, -273, -271, -263, -263, -256, -251, -243, -229, -223, -212, -207, -201, -199, -199, -198, -186, -175, -162, -161, -161, -146, -140, -138, -135, -127, -125, -113, -96, -90, -85, -77, -75, -60, -52, -26, -11, -7, 0, 0, 5, 10, 29, 46, 67, 67, 68, 70, 92, 106, 110, 138, 147, 149, 155, 159, 178, 178, 191, 196, 197, 200, 210, 211, 226, 233, 239, 248, 258, 265, 267, 271, 285, 288, 293, 311, 318, 320, 324, 325, 327, 333, 340, 374, 375, 379, 386, 390, 393, 403, 410, 415, 418, 419, 462, 482, 486, 486, 489, 493, 501, 508, 524, 527, 532, 532, 534, 557, 568, 573, 580, 598, 603, 604, 608, 609, 620, 625, 626, 630, 663, 664, 698, 706, 716, 727, 739, 751, 752, 759, 761, 781, 785, 807, 807, 816, 820, 820, 841, 844, 849, 856, 859, 877, 877, 882, 889, 890, 917, 918, 920, 923, 926, 929, 937, 950, 956, 956, 971, 974, 978, 989, 998, 999, 1005, 1016, 1022, 1056, 1058, 1059, 1068, 1082, 1085, 1085, 1091, 1097, 1104, 1117, 1124, 1126, 1130, 1134, 1136, 1143, 1152, 1173, 1192, 1197, 1203, 1203, 1216, 1217, 1235, 1248, 1268, 1272, 1275, 1277, 1280, 1286, 1305, 1314, 1340, 1356, 1361, 1363, 1368, 1394, 1398, 1402, 1404, 1423, 1428, 1431, 1435, 1456, 1457, 1468, 1475, 1484, 1498, 1498, 1534, 1542, 1553, 1570, 1571, 1590, 1594, 1599, 1611, 1625, 1633, 1649, 1666, 1670, 1684, 1687, 1688, 1701, 1706, 1709, 1712, 1712, 1726, 1732, 1744, 1753, 1755, 1775, 1780, 1786, 1791, 1804, 1823, 1834, 1847, 1852, 1865, 1865, 1878, 1883, 1890, 1896, 1908, 1915, 1927, 1932, 1936, 1948, 1950, 1968, 1968, 1990, 1991, 1991},
	}
	for _, c := range cases {
		fmt.Printf("max count: %+v, %d\n\n", c, maximumCount(c))
	}
}
