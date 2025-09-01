package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetFiles() {
	files := []*m.File{
		// Class Backgrounds
		{
			ID:           1,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      1,
			Filename:     "SUJGY2ic8kXZPYGhpiJCUbvk1TqlLONo.png",
			FileLocation: "classes/1",
		},
		{
			ID:           2,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      2,
			Filename:     "5MoE3PczPvt79oliOTZ9PTwTatrpsY1i.png",
			FileLocation: "classes/2",
		},
		{
			ID:           3,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      3,
			Filename:     "8TlHL6RQM6a97DjmHu1pcWlzIwcpRd8o.png",
			FileLocation: "classes/3",
		},
		{
			ID:           4,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      4,
			Filename:     "9z53I0sM3xyYTUgvruDQOSWpi52lZMt5.png",
			FileLocation: "classes/4",
		},
		{
			ID:           5,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      5,
			Filename:     "l9ZN3QdXo3MucoD9QTO0vs0VuaAZTAZ3.png",
			FileLocation: "classes/5",
		},
		{
			ID:           6,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      6,
			Filename:     "HCwMvRTCYievZOFnm4dXC9D4f8DuMwir.png",
			FileLocation: "classes/6",
		},
		{
			ID:           7,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      7,
			Filename:     "3uqED6MCf7mYxtUBfzWzBM4nRNWpQLVh.png",
			FileLocation: "classes/7",
		},
		{
			ID:           8,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      8,
			Filename:     "ZcEf3rcJgaSFLooRJU1Ape4afdNh9y99.png",
			FileLocation: "classes/8",
		},
		{
			ID:           9,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      9,
			Filename:     "AhjIUsuF0OYA1DOJIoKOCcQ0fF3tEKzg.png",
			FileLocation: "classes/9",
		},
		{
			ID:           10,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      10,
			Filename:     "nlR8FoIT1QPDFy32lbjZPShlEvj8RYfZ.png",
			FileLocation: "classes/10",
		},
		{
			ID:           11,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      11,
			Filename:     "M3owM13vmx4uStXbBU4fWw5thAGwijli.png",
			FileLocation: "classes/11",
		},
		{
			ID:           12,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      12,
			Filename:     "eqmjb4U9QxCIJU19j5akQDfOuW3jl6qa.png",
			FileLocation: "classes/12",
		},
		{
			ID:           13,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      13,
			Filename:     "W8RnfwEMhOhpFq1aZGaye7DKqPNyFGnW.png",
			FileLocation: "classes/13",
		},
		{
			ID:           14,
			Model:        m.FileModelClassBackgroundImage,
			ModelId:      14,
			Filename:     "Y7yCA4l0CyUHSTzzTAXGopdPwVH3JrDx.png",
			FileLocation: "classes/14",
		},
		// Class Logos
		{
			ID:           15,
			Model:        m.FileModelClassLogo,
			ModelId:      1,
			Filename:     "b4W5kqTgj59jqxUiWEIagpaxIFNXidQF.jpeg",
			FileLocation: "classes/1",
		},
		{
			ID:           16,
			Model:        m.FileModelClassLogo,
			ModelId:      2,
			Filename:     "oXQuyAKM3TdRArzvS277GwzARneye3pn.jpeg",
			FileLocation: "classes/2",
		},
		{
			ID:           17,
			Model:        m.FileModelClassLogo,
			ModelId:      3,
			Filename:     "DDy6f8aB3OLlc4xKMZrk09kIgfCnZ0DH.jpeg",
			FileLocation: "classes/3",
		},
		{
			ID:           18,
			Model:        m.FileModelClassLogo,
			ModelId:      4,
			Filename:     "LcnwDmJZfdjF6wRFr9yjvOK2SYJWu9oO.jpeg",
			FileLocation: "classes/4",
		},
		{
			ID:           19,
			Model:        m.FileModelClassLogo,
			ModelId:      5,
			Filename:     "HAgiAVUr2BLEK01lkwI5MYwJS99uh5j4.jpeg",
			FileLocation: "classes/5",
		},
		{
			ID:           20,
			Model:        m.FileModelClassLogo,
			ModelId:      6,
			Filename:     "VHAILdFL3MsEmTY9jdgRaC1sWuohkmzk.jpeg",
			FileLocation: "classes/6",
		},
		{
			ID:           21,
			Model:        m.FileModelClassLogo,
			ModelId:      7,
			Filename:     "nEvF9BZMdpYUYWQAonx1RzbF3OwBv89M.jpeg",
			FileLocation: "classes/7",
		},
		{
			ID:           22,
			Model:        m.FileModelClassLogo,
			ModelId:      8,
			Filename:     "rjz5gDlhSKrgDtEtgeDcmZsOngvnbg39.jpeg",
			FileLocation: "classes/8",
		},
		{
			ID:           23,
			Model:        m.FileModelClassLogo,
			ModelId:      9,
			Filename:     "XzhR9MvJjPfn515E1ij2bLzCrJqqJ29H.jpeg",
			FileLocation: "classes/9",
		},
		{
			ID:           24,
			Model:        m.FileModelClassLogo,
			ModelId:      10,
			Filename:     "Vknvnnco1MU8fGyz46tIG8TO2sIkQEVI.jpeg",
			FileLocation: "classes/10",
		},
		{
			ID:           25,
			Model:        m.FileModelClassLogo,
			ModelId:      11,
			Filename:     "zC7JkZJbcMHdgzeGOZOqBKy7dRp23gpS.jpeg",
			FileLocation: "classes/11",
		},
		{
			ID:           26,
			Model:        m.FileModelClassLogo,
			ModelId:      12,
			Filename:     "rtUxRMYkKKqecjBi8wm0m2RuKnOshMfk.jpeg",
			FileLocation: "classes/12",
		},
		{
			ID:           27,
			Model:        m.FileModelClassLogo,
			ModelId:      13,
			Filename:     "YVjzL8LiakLhcrPvwxZduZWTXvsCFani.png",
			FileLocation: "classes/13",
		},
		{
			ID:           28,
			Model:        m.FileModelClassLogo,
			ModelId:      14,
			Filename:     "f8iFRyAg7kZV8IXoxniy1LQGTCjDG6IV.jpeg",
			FileLocation: "classes/14",
		},
		// Race Logos
		{
			ID:           29,
			Model:        m.FileModelRaceLogo,
			ModelId:      1,
			Filename:     "2ABX7dxgzVoyos02Jbjhn9RTld049AKh.jpeg",
			FileLocation: "races/1",
		},
		{
			ID:           30,
			Model:        m.FileModelRaceLogo,
			ModelId:      2,
			Filename:     "fKx6AEPNtHXoFLncCaaUflDMEaioAlbN.jpeg",
			FileLocation: "races/2",
		},
		{
			ID:           31,
			Model:        m.FileModelRaceLogo,
			ModelId:      3,
			Filename:     "YwmEMffB7KjMKrvWjpOWxS9yVyoTFbAl.jpeg",
			FileLocation: "races/3",
		},
		{
			ID:           32,
			Model:        m.FileModelRaceLogo,
			ModelId:      4,
			Filename:     "1Mii5gkrXFWGvhBsKJELzS6fHujUknjc.jpeg",
			FileLocation: "races/4",
		},
		{
			ID:           33,
			Model:        m.FileModelRaceLogo,
			ModelId:      5,
			Filename:     "PUZr5nt9Jqmi0kaqL3X782FtGVhnHc1u.png",
			FileLocation: "races/5",
		},
		{
			ID:           34,
			Model:        m.FileModelRaceLogo,
			ModelId:      6,
			Filename:     "RFKObdvvVkNZbliTVSnmzmJWuA0Smg7W.jpeg",
			FileLocation: "races/6",
		},
		{
			ID:           35,
			Model:        m.FileModelRaceLogo,
			ModelId:      7,
			Filename:     "iwTJvPXu9ZL634xhgYwTtNn2ywrPkXIy.jpeg",
			FileLocation: "races/7",
		},
		{
			ID:           36,
			Model:        m.FileModelRaceLogo,
			ModelId:      8,
			Filename:     "yORA72xGPiFshkSoQvNctcHCyrIXdZNu.png",
			FileLocation: "races/8",
		},
		{
			ID:           37,
			Model:        m.FileModelRaceLogo,
			ModelId:      9,
			Filename:     "pzYR1IcbzRBY4wV1ULWQW4VlUxe8PIBW.jpeg",
			FileLocation: "races/9",
		},
		{
			ID:           38,
			Model:        m.FileModelRaceLogo,
			ModelId:      10,
			Filename:     "KdBWefFAnvZ5gJf0F4XrK2mdB6sj7LrS.jpeg",
			FileLocation: "races/10",
		},
		{
			ID:           39,
			Model:        m.FileModelRaceLogo,
			ModelId:      11,
			Filename:     "58IxFOBkhJwI1K0QgqQuPsLEaBMNWJid.jpeg",
			FileLocation: "races/11",
		},
		{
			ID:           40,
			Model:        m.FileModelRaceLogo,
			ModelId:      12,
			Filename:     "plvxAguXvUKbUIuKZhPxR8K25LYDMCs7.jpeg",
			FileLocation: "races/12",
		},
		{
			ID:           41,
			Model:        m.FileModelRaceLogo,
			ModelId:      13,
			Filename:     "82ydTG2WRFdC5jPCtwVTfw5hoJZc8XPf.jpeg",
			FileLocation: "races/13",
		},
		{
			ID:           42,
			Model:        m.FileModelRaceLogo,
			ModelId:      14,
			Filename:     "iB3JIatoszWRu1uTTyGSIbvViru1g6cx.jpeg",
			FileLocation: "races/14",
		},
		{
			ID:           43,
			Model:        m.FileModelRaceLogo,
			ModelId:      15,
			Filename:     "F3mNNcwA3kDoHZcGNrY9qQeWgTxV2sft.jpeg",
			FileLocation: "races/15",
		},
		// Character Profile Pictures
		{
			ID:           44,
			Model:        m.FileModelCharacterProfilePicture,
			ModelId:      1,
			Filename:     "3lGxLzrSYAwGxtWedqYmy0nLVKKdpOnP.png",
			FileLocation: "characters/1",
		},
	}

	for _, file := range files {
		err := s.DB.Where("id = ?", file.ID).FirstOrCreate(&file).Error
		if err != nil {
			log.Printf("Error creating file with id %v in seeder: %v", file.ID, err.Error())
		}
	}
}
