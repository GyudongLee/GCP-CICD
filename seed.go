package main

var seed = []Exoplanet{
	{
		Name:   "2M 0746+20 b",
		Mass:   30,
		Radius: 0.97,
		Period: 4640,
	},
	{
		Name:   "2M 2140+16 b",
		Mass:   20,
		Radius: 0.92,
		Period: 7340,
	},
	{
		Name:   "2M 2206-20 b",
		Mass:   30,
		Radius: 1.3,
		Period: 8686,
	},
	{
		Name:   "51 Eri b",
		Mass:   2,
		Radius: 1,
		Period: 15000,
	},
	{
		Name:   "55 Cancri e",
		Mass:   0.0251,
		Radius: 0.16728,
		Period: 0.7365474,
	},
	{
		Name:   "BD+20 594 b",
		Mass:   0.05129,
		Radius: 0.1989,
		Period: 41.6855,
	},
	{
		Name:   "beta Pic b",
		Mass:   11,
		Radius: 1.65,
		Period: 8167,
	},
	{
		Name:   "CoRoT-10 b",
		Mass:   2.75,
		Radius: 0.97,
		Period: 13.2406,
	},
	{
		Name:   "CoRoT-11 b",
		Mass:   2.49,
		Radius: 1.39,
		Period: 2.994325,
	},
	{
		Name:   "CoRoT-12 b",
		Mass:   0.917,
		Radius: 1.44,
		Period: 2.828042,
	},
	{
		Name:   "CoRoT-13 b",
		Mass:   1.308,
		Radius: 0.885,
		Period: 4.03519,
	},
	{
		Name:   "CoRoT-14 b",
		Mass:   7.6,
		Radius: 1.09,
		Period: 1.51214,
	},
	{
		Name:   "CoRoT-16 b",
		Mass:   0.53,
		Radius: 1.17,
		Period: 5.35227,
	},
	{
		Name:   "CoRoT-17 b",
		Mass:   2.45,
		Radius: 1.02,
		Period: 3.768125,
	},
	{
		Name:   "CoRoT-18 b",
		Mass:   3.47,
		Radius: 1.31,
		Period: 1.9000693,
	},
	{
		Name:   "CoRoT-19 b",
		Mass:   1.11,
		Radius: 1.45,
		Period: 3.89713,
	},
	{
		Name:   "CoRoT-1 b",
		Mass:   1.03,
		Radius: 1.49,
		Period: 1.5089557,
	},
	{
		Name:   "CoRoT-20 b",
		Mass:   4.3,
		Radius: 0.84,
		Period: 9.24285,
	},
	{
		Name:   "CoRoT-21 b",
		Mass:   2.26,
		Radius: 1.3,
		Period: 2.72474,
	},
	{
		Name:   "CoRoT-22 b",
		Mass:   0.0383853003178,
		Radius: 0.435365216104,
		Period: 9.75598,
	},
	{
		Name:   "CoRoT-23 b",
		Mass:   2.8,
		Radius: 1.05,
		Period: 3.6314,
	},
	{
		Name:   "CoRoT-24 c",
		Mass:   0.088,
		Radius: 0.44,
		Period: 11.759,
	},
	{
		Name:   "CoRoT-25 b",
		Mass:   0.27,
		Radius: 1.08,
		Period: 4.86069,
	},
	{
		Name:   "CoRoT-26 b",
		Mass:   0.52,
		Radius: 1.26,
		Period: 4.20474,
	},
	{
		Name:   "CoRoT-27 b",
		Mass:   10.39,
		Radius: 1.007,
		Period: 3.57532,
	},
	{
		Name:   "CoRoT-28 b",
		Mass:   0.484,
		Radius: 0.955,
		Period: 5.20851,
	},
	{
		Name:   "CoRoT-29 b",
		Mass:   0.85,
		Radius: 0.97,
		Period: 2.8505615,
	},
	{
		Name:   "CoRoT-2 b",
		Mass:   3.31,
		Radius: 1.465,
		Period: 1.7429964,
	},
	{
		Name:   "CoRoT-30 b",
		Mass:   2.84,
		Radius: 1.02,
		Period: 9.06005,
	},
	{
		Name:   "CoRoT-31 b",
		Mass:   0.85,
		Radius: 1.5,
		Period: 4.62941,
	},
	{
		Name:   "CoRoT-3 b",
		Mass:   21.66,
		Radius: 1.01,
		Period: 4.25680,
	},
	{
		Name:   "CoRoT-4 b",
		Mass:   0.72,
		Radius: 1.19,
		Period: 9.20205,
	},
	{
		Name:   "CoRoT-5 b",
		Mass:   0.467,
		Radius: 1.388,
		Period: 4.0378962,
	},
	{
		Name:   "CoRoT-6 b",
		Mass:   2.96,
		Radius: 1.166,
		Period: 8.886593,
	},
	{
		Name:   "CoRoT-7 b",
		Mass:   0.0151,
		Radius: 0.15,
		Period: 0.853585,
	},
	{
		Name:   "CoRoT-8 b",
		Mass:   0.22,
		Radius: 0.57,
		Period: 6.21229,
	},
	{
		Name:   "CoRoT-9 b",
		Mass:   0.84,
		Radius: 1.05,
		Period: 95.2738,
	},
	{
		Name:   "EPIC 201367065 b",
		Mass:   0.0264,
		Radius: 0.194,
		Period: 10.05449,
	},
	{
		Name:   "EPIC 201505350 c",
		Mass:   0.0500,
		Radius: 0.434,
		Period: 11.90715,
	},
	{
		Name:   "EPIC 201505350 b",
		Mass:   0.138,
		Radius: 0.691,
		Period: 7.91940,
	},
	{
		Name:   "EPIC 201546283 b",
		Mass:   0.09156,
		Radius: 0.3970,
		Period: 6.77145,
	},
	{
		Name:   "EPIC 201577035 b",
		Mass:   0.0850,
		Radius: 0.3426,
		Period: 19.3044,
	},
	{
		Name:   "EPIC 201912552 b",
		Mass:   0.02504,
		Radius: 0.2123,
		Period: 32.93963,
	},
	{
		Name:   "EPIC 203771098 b",
		Mass:   0.066,
		Radius: 0.520,
		Period: 20.88508,
	},
	{
		Name:   "EPIC 203771098 c",
		Mass:   0.0850,
		Radius: 0.723,
		Period: 42.36342,
	},
	{
		Name:   "EPIC 204221263 b",
		Mass:   0.0378,
		Radius: 0.138,
		Period: 4.01593,
	},
	{
		Name:   "EPIC 204221263 c",
		Mass:   0.031,
		Radius: 0.216,
		Period: 10.56103,
	},
	{
		Name:   "EPIC 205071984 b",
		Mass:   0.06639,
		Radius: 0.480,
		Period: 8.99218,
	},
	{
		Name:   "K2-33 b",
		Mass:   3.6,
		Radius: 0.451,
		Period: 5.42513,
	},
	{
		Name:   "EPIC 210894022 b",
		Mass:   0.027,
		Radius: 0.17,
		Period: 5.35117,
	},
	{
		Name:   "K2-29 b",
		Mass:   0.73,
		Radius: 1.19,
		Period: 3.2588321,
	},
}
