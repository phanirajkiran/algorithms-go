package _1339_maximum_product_of_splitted_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

var bigInput1 = []int{3434, 4223, 2441, 6764, 5911, 7094, 1827, 9223, 3580, 6615, 8446, 2770, 5112, 718, 3292, 4092, 3269, 377, 7407, 4515, 4512, 6098, 282, 2197, 9833, 5285, 5841, 9643, 8708, 500, 1834, 7466, 1360, 8075, 9353, 804, 656, 8645, 2445, 4648, 1194, 2185, 7883, btree.Null, 3282, 2067, 8329, 4847, 1363, 1037, 2829, 3789, 1321, 8183, 2392, 8978, 436, 7776, 2286, 8635, 587, 4391, 5075, 7307, 8431, 2236, 3588, btree.Null, btree.Null, 6968, 6324, btree.Null, 2149, btree.Null, 5868, 7401, btree.Null, 8175, 7064, 1404, 8772, btree.Null, btree.Null, 2259, 3610, 2455, 7961, btree.Null, 3397, 8996, 7112, 1316, 4197, 8704, 2391, 227, 4720, 9266, 3273, 3503, btree.Null, 5237, 7905, 1921, 8540, 1886, 6681, 4740, 6134, 8408, 3442, btree.Null, 3830, 2786, 5382, 3499, 4469, 1260, 1456, 6568, 746, 8076, 1665, btree.Null, 5700, 7959, 209, 2485, btree.Null, 3253, 6181, 1080, 8731, 4829, 7285, btree.Null, 2136, 3995, 3153, 4968, 549, 3290, 627, 7812, 4406, 254, 8382, btree.Null, btree.Null, btree.Null, btree.Null, 4246, 5958, 8358, 1853, 2260, 3188, 1963, 9753, btree.Null, 8976, btree.Null, 1244, btree.Null, btree.Null, 9473, 8385, btree.Null, 2370, 3469, 2059, 9616, 1238, 3089, 9857, btree.Null, 3873, 2465, 1945, 6202, 7906, 9853, 2006, btree.Null, btree.Null, btree.Null, 9707, 4539, 4815, 3158, 4493, 5941, 7648, 7964, 6534, 9843, 7333, 409, 5246, btree.Null, btree.Null, 7899, btree.Null, 2515, 5559, 2888, btree.Null, 9258, 9370, 9767, btree.Null, 4140, 9421, 5873, 6398, 1152, 6895, 9101, 2652, 6053, 721, btree.Null, btree.Null, btree.Null, 303, 4819, 1605, 3110, 5366, 5363, 6173, 6188, 7605, 6414, btree.Null, 1038, 3255, 870, 798, 217, 1880, 76, 726, 8296, 4856, btree.Null, 581, btree.Null, btree.Null, 3247, 261, 7183, 5597, 6299, btree.Null, 6591, btree.Null, 2059, 4841, 3925, btree.Null, btree.Null, btree.Null, 8182, 3954, 4280, 5532, btree.Null, btree.Null, 1900, 1084, 1704, 459, 4312, btree.Null, btree.Null, 5370, 6139, 5517, 1806, 2949, 3257, btree.Null, 7820, 7347, btree.Null, btree.Null, 1224, 1430, 9544, 1349, 3583, 5668, btree.Null, 2646, 3945, 1422, 9511, btree.Null, btree.Null, 5411, 8568, btree.Null, 5311, 7155, 9720, 1904, 3772, btree.Null, 3186, btree.Null, btree.Null, 9768, 2530, btree.Null, btree.Null, btree.Null, 7092, btree.Null, 7664, btree.Null, 6784, 2692, 2031, 9754, 9864, 4004, btree.Null, 4518, 4412, 8741, 5750, 9149, btree.Null, 6613, 1859, btree.Null, 784, 2634, 6172, 4635, btree.Null, btree.Null, 2675, btree.Null, btree.Null, 718, 5809, 584, btree.Null, 4540, 1547, 5399, btree.Null, 2446, btree.Null, 3829, btree.Null, btree.Null, btree.Null, 8734, 6278, 8937, 5207, 2497, btree.Null, 4524, 502, 9879, 778, 409, 1892, btree.Null, 7085, 9924, 2346, btree.Null, 8991, btree.Null, btree.Null, 9519, btree.Null, btree.Null, 384, btree.Null, 3435, 5565, btree.Null, 4053, 7290, btree.Null, btree.Null, btree.Null, 4105, 2830, 5815, 8548, btree.Null, btree.Null, 9641, 8344, 3660, btree.Null, 9721, 115, 8064, 3557, 7344, 6233, btree.Null, 6000, 9789, 8861, 3678, 1170, 1082, 6525, 6463, 9355, btree.Null, btree.Null, btree.Null, 1227, 7890, 9396, btree.Null, btree.Null, 2978, 7651, btree.Null, btree.Null, btree.Null, btree.Null, 7948, 3720, 3218, 1878, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3308, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 201, 4360, 3255, btree.Null, btree.Null, 2568, 3229, btree.Null, 3261, 6406, btree.Null, 5929, 9516, btree.Null, btree.Null, btree.Null, 9247, 5988, 2073, btree.Null, btree.Null, btree.Null, 8689, btree.Null, btree.Null, btree.Null, btree.Null, 5819, btree.Null, btree.Null, btree.Null, 1352, btree.Null, 9355, 1350, btree.Null, btree.Null, 9825, 6797, 1522, btree.Null, 9138, btree.Null, btree.Null, 1040, btree.Null, btree.Null, 8289, 6770, 6913, 7863, 5710, 6803, btree.Null, 2297, btree.Null, btree.Null, btree.Null, 9738, btree.Null, btree.Null, 1004, btree.Null, btree.Null, btree.Null, btree.Null, 4112, 2947, 913, 8498, 8879, btree.Null, btree.Null, 8526, 2609, 3486, 7499, btree.Null, btree.Null, 9878, 1833, 309, 5298, 6989, 3188, 8401, 7437, 2560, 831, btree.Null, 4949, btree.Null, btree.Null, 5723, 8576, 1991, btree.Null, 7982, 4985, 4376, 4170, 5564, 5641, btree.Null, btree.Null, 1139, 2392, btree.Null, btree.Null, btree.Null, 4040, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9909, btree.Null, btree.Null, 2298, btree.Null, btree.Null, 2294, 496, 9336, 304, 8230, btree.Null, 6056, btree.Null, btree.Null, 5382, 759, 7933, 5960, 3873, 4548, 6075, 2686, 7177, btree.Null, 4091, btree.Null, 767, btree.Null, 4877, 1756, 1717, 4537, btree.Null, btree.Null, btree.Null, btree.Null, 2095, 7406, 870, btree.Null, btree.Null, btree.Null, 1509, 8381, btree.Null, 3923, 5682, 667, 4522, btree.Null, 1869, 7024, 9323, 1579, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 2951, 3890, btree.Null, 2131, btree.Null, 8785, 757, btree.Null, 5912, 2230, 7837, 4311, 1439, 6591, 5966, btree.Null, 2452, 7591, 2279, 9438, 8189, 2293, 8306, 2934, 3316, 7320, 4994, 4027, btree.Null, 6188, 2554, btree.Null, 9148, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 1252, 5554, 7583, 1552, 3862, btree.Null, 8131, btree.Null, 745, btree.Null, btree.Null, 8880, 16, 7642, 7181, btree.Null, btree.Null, 2739, btree.Null, 4819, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 1490, btree.Null, 9057, btree.Null, btree.Null, btree.Null, 1775, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3482, btree.Null, 7145, btree.Null, 2374, 9308, btree.Null, btree.Null, 8204, btree.Null, btree.Null, btree.Null, 928, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3209, 6076, btree.Null, btree.Null, 9683, btree.Null, 4879, btree.Null, 7874, btree.Null, 4504, 69, btree.Null, btree.Null, btree.Null, 1, 362, btree.Null, 1404, 8888, 1875, 6483, 6565, 8395, 2214, 8420, 4913, 7525, 7657, 327, 5459, 2016, 6299, 5042, 8558, 876, 4798, btree.Null, 7907, 4536, btree.Null, 4508, btree.Null, btree.Null, 814, 6337, btree.Null, 2782, 3291, 5829, btree.Null, 7395, 9107, 9361, 1260, 10000, 769, 8659, btree.Null, btree.Null, 9771, btree.Null, 1160, 5252, 4088, 3967, 2779, 7333, btree.Null, btree.Null, 872, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8543, btree.Null, btree.Null, 9733, 5123, 4585, btree.Null, 2880, 9590, btree.Null, btree.Null, btree.Null, 5920, btree.Null, 5902, 7266, btree.Null, 9077, btree.Null, btree.Null, 3794, btree.Null, 2867, btree.Null, 1822, btree.Null, btree.Null, 3104, 4782, 9380, 3245, btree.Null, 8841, 29, btree.Null, btree.Null, 4115, 4113, 9770, 4838, btree.Null, btree.Null, btree.Null, 4573, btree.Null, btree.Null, btree.Null, 5290, btree.Null, btree.Null, 8254, btree.Null, btree.Null, 5139, 6, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9066, btree.Null, 4679, 1378, 1120, 6088, 2864, btree.Null, btree.Null, 4555, btree.Null, btree.Null, 4250, 9565, 7891, 8402, 3625, btree.Null, 2381, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 980, btree.Null, 5230, 6426, 5967, 750, btree.Null, btree.Null, 4537, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3030, btree.Null, btree.Null, 864, btree.Null, 2600, btree.Null, 8858, 1373, 2775, 6504, btree.Null, 5279, btree.Null, btree.Null, btree.Null, 3933, 9352, 8292, 1023, btree.Null, btree.Null, 5354, 3052, btree.Null, 6710, btree.Null, 2240, btree.Null, btree.Null, 2165, btree.Null, btree.Null, btree.Null, 2958, btree.Null, 4988, btree.Null, 5882, btree.Null, btree.Null, btree.Null, 8686, 3471, 244, 9758, 96, btree.Null, 3855, btree.Null, btree.Null, btree.Null, btree.Null, 3210, btree.Null, 8274, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 4513, btree.Null, btree.Null, 5767, btree.Null, btree.Null, btree.Null, btree.Null, 9506, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 6745, 5337, btree.Null, 3575, 9091, btree.Null, btree.Null, btree.Null, 9830, btree.Null, 9738, 7999, btree.Null, 6138, btree.Null, 9425, 4203, 3839, 3824, btree.Null, 9591, 9675, btree.Null, 6884, 56, 2355, 5197, 8046, 4136, 6357, btree.Null, 8317, 4117, btree.Null, 1842, 7985, 6869, btree.Null, btree.Null, 9812, 8336, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 7511, 1385, btree.Null, 2986, btree.Null, 8634, 308, 4008, 7522, btree.Null, btree.Null, btree.Null, btree.Null, 5879, btree.Null, btree.Null, 2858, btree.Null, 4856, 5299, 8482, 5326, 2624, 858, 4350, 7321, 521, 8806, 6941, 3319, btree.Null, 8915, 1065, 3453, 6126, 5694, btree.Null, btree.Null, btree.Null, btree.Null, 9768, 123, 6371, btree.Null, btree.Null, 2088, 5298, 782, 9818, btree.Null, 8197, 1175, 1146, 1945, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 4079, btree.Null, btree.Null, btree.Null, 3231, btree.Null, btree.Null, btree.Null, btree.Null, 8841, 8371, 8971, btree.Null, 1755, 1891, btree.Null, btree.Null, btree.Null, btree.Null, 3016, btree.Null, btree.Null, 3136, btree.Null, 678, btree.Null, 7919, 644, btree.Null, btree.Null, btree.Null, btree.Null, 8769, btree.Null, 3159, btree.Null, btree.Null, btree.Null, 6179, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 2913, 7271, btree.Null, 733, 3795, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 7164, 2469, 2744, 5399, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8888, 7666, btree.Null, 6396, 1148, 2642, btree.Null, 4113, 555, 1669, btree.Null, btree.Null, btree.Null, btree.Null, 3039, 9576, 7126, btree.Null, 263, btree.Null, btree.Null, btree.Null, btree.Null, 9748, btree.Null, btree.Null, 9176, 5235, 222, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 6188, btree.Null, btree.Null, btree.Null, 6012, 8488, btree.Null, btree.Null, btree.Null, btree.Null, 888, btree.Null, btree.Null, 2384, btree.Null, 6752, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8105, 1530, 8038, btree.Null, btree.Null, btree.Null, 9791, 5982, 4757, 2791, 9973, btree.Null, 1030, btree.Null, btree.Null, btree.Null, 5403, 2796, btree.Null, btree.Null, 4324, 4362, 1124, 7488, 4436, 2556, 5331, btree.Null, 8660, 1594, 7398, 8895, 9023, btree.Null, btree.Null, 577, 5560, 5592, btree.Null, 3838, btree.Null, 336, btree.Null, btree.Null, btree.Null, btree.Null, 5396, btree.Null, 5866, btree.Null, btree.Null, btree.Null, 2308, btree.Null, btree.Null, 9077, btree.Null, btree.Null, 1270, btree.Null, 1593, 1990, btree.Null, btree.Null, 6756, btree.Null, btree.Null, 9650, btree.Null, btree.Null, 5707, 1106, btree.Null, 5615, 6076, 825, btree.Null, btree.Null, 6629, btree.Null, btree.Null, btree.Null, btree.Null, 1502, 4739, btree.Null, btree.Null, 688, btree.Null, 5134, 4677, btree.Null, btree.Null, btree.Null, btree.Null, 9885, 1152, btree.Null, btree.Null, btree.Null, btree.Null, 1619, 8418, 7699, 8855, btree.Null, btree.Null, 8402, 9539, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9468, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 5624, btree.Null, 9683, btree.Null, btree.Null, btree.Null, 2742, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8847, 7327, btree.Null, 2759, 2537, btree.Null, btree.Null, btree.Null, 1981, btree.Null, btree.Null, 4756, 5394, 1265, 2611, btree.Null, 4864, 7675, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 213, btree.Null, btree.Null, btree.Null, 151, btree.Null, btree.Null, btree.Null, 7739, 5814, 130, btree.Null, btree.Null, btree.Null, 481, 1623, 4669, btree.Null, 8861, 9953, 5835, 1593, 4338, 8037, 2690, 7015, 6100, 652, 1497, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 6487, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 4692, 2297, btree.Null, 5542, btree.Null, btree.Null, btree.Null, btree.Null, 9787, btree.Null, btree.Null, btree.Null, 149, 2236, 4955, btree.Null, btree.Null, btree.Null, btree.Null, 6703, btree.Null, 5427, 4698, 2349, 4578, 9504, btree.Null, btree.Null, btree.Null, btree.Null, 9921, 473, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 2242, 6092, 4113, 6763, btree.Null, btree.Null, 8592, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9095, btree.Null, btree.Null, btree.Null, 9075, btree.Null, btree.Null, btree.Null, btree.Null, 8514, btree.Null, btree.Null, 5463, 4782, 2528, 1399, btree.Null, 3188, btree.Null, btree.Null, 8314, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3321, 1832, 8620, 4797, 9036, btree.Null, 1243, btree.Null, btree.Null, btree.Null, 7464, btree.Null, btree.Null, 9351, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 7468, 425, btree.Null, 3747, btree.Null, btree.Null, btree.Null, 7599, btree.Null, 3724, btree.Null, 742, 9578, 6532, 2358, 8082, 340, 9320, 9391, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 5069, btree.Null, btree.Null, btree.Null, 3669, btree.Null, 598, 629, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 2557, btree.Null, btree.Null, btree.Null, btree.Null, 7496, 6361, btree.Null, 3693, btree.Null, 1769, 5807, 4896, 3384, 2709, 8316, 3453, btree.Null, 8683, 6104, 1462, btree.Null, 3274, 5488, btree.Null, 9127, 1830, btree.Null, btree.Null, btree.Null, 4329, btree.Null, 8761, btree.Null, btree.Null, btree.Null, btree.Null, 1267, btree.Null, 6006, 2603, btree.Null, 8455, btree.Null, btree.Null, 312, btree.Null, 5447, 9453, btree.Null, 7068, 5531, 4462, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 4103, btree.Null, btree.Null, 3212, 2659, btree.Null, btree.Null, btree.Null, btree.Null, 1520, 1032, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 1495, 3098, 709, 9217, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8529, 4539, 2454, 8339, 2077, 6314, 5155, 7575, 5197, 7167, btree.Null, btree.Null, btree.Null, 9379, btree.Null, btree.Null, btree.Null, 9393, 8408, btree.Null, 8633, btree.Null, btree.Null, btree.Null, btree.Null, 6241, 8583, btree.Null, 82, btree.Null, btree.Null, btree.Null, 429, 9391, 6592, 418, 6105, 5964, btree.Null, 9746, 8472, btree.Null, btree.Null, btree.Null, 5501, btree.Null, 3570, btree.Null, btree.Null, btree.Null, btree.Null, 9490, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3962, btree.Null, btree.Null, 2299, 7952, 356, btree.Null, btree.Null, btree.Null, 5180, 8541, 1860, btree.Null, btree.Null, btree.Null, btree.Null, 4862, btree.Null, 1825, btree.Null, 3981, btree.Null, btree.Null, btree.Null, 5022, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 6694, btree.Null, btree.Null, 5648, 2911, 5099, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8719, 3130, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9453, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 5808, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9101, 877, 5009, 1889, 9232, 3699, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 2971, 6579, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3152, btree.Null, btree.Null, 8947, btree.Null, 7526, btree.Null, btree.Null, 8076, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 4827, 4544, btree.Null, btree.Null, btree.Null, btree.Null, 1373, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9681, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 3068, btree.Null, 4929, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 1447, btree.Null, 171, btree.Null, 4767, btree.Null, btree.Null, btree.Null, 7197, btree.Null, btree.Null, 3116, 2056, 5742, btree.Null, 7142, 7645, 692, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 7244, btree.Null, btree.Null, btree.Null, btree.Null, 9336, 9541, 1885, btree.Null, btree.Null, btree.Null, btree.Null, 1095, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8997, btree.Null, btree.Null, 3604, 4473, btree.Null, btree.Null, btree.Null, 7284, 298, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 9330, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 5501, btree.Null, btree.Null, btree.Null, btree.Null, 2950, 5479, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 8774, 5861, btree.Null, 2043}
