package forge

import (
	"github.com/mcuadros/go-version"
)

type Number struct {
	ID        string `json:"version"`
	Minecraft string `json:"mcversion"`
}

func (s *Number) Invalid() bool {
	return version.Compare(s.Minecraft, "1.6.4", "<")
}

// http://files.minecraftforge.net/maven/net/minecraftforge/forge/1.8.9-11.15.1.1764/forge-1.8.9-11.15.1.1764-installer.jar
// http://files.minecraftforge.net/maven/net/minecraftforge/forge/1.8.9-11.15.1.1764/forge-1.8.9-11.15.1.1764-universal.jar

// http://files.minecraftforge.net/maven/net/minecraftforge/forge/1.7.10-10.13.4.1614-1.7.10/forge-1.7.10-10.13.4.1614-1.7.10-installer.jar
// http://files.minecraftforge.net/maven/net/minecraftforge/forge/1.7.10-10.13.4.1614-1.7.10/forge-1.7.10-10.13.4.1614-1.7.10-universal.jar

// http://files.minecraftforge.net/maven/net/minecraftforge/forge/1.1-1.3.4.29/forge-1.1-1.3.4.29-server.zip

// {
//   "homepage": "http://files.minecraftforge.net/maven/net/minecraftforge/forge/",
//   "webpath": "http://files.minecraftforge.net/maven/net/minecraftforge/forge/"
//   "number": {
//     "1": {
//       "branch": null,
//       "build": 1,
//       "files": [
//         [
//           "zip",
//           "src",
//           "fd397591148fac49a7d57aafdccac6a3"
//         ],
//         [
//           "zip",
//           "client",
//           "4d96d6e8f1543c5fa1184f4771ce16e2"
//         ],
//         [
//           "zip",
//           "server",
//           "d7e1df9a91ded33be81ee8226b027c2f"
//         ],
//         [
//           "txt",
//           "changelog",
//           "df98aec1a868ce99532c64f246387d55"
//         ]
//       ],
//       "mcversion": "1.1",
//       "modified": 1412001430.0,
//       "version": "1.3.2.1"
//     },
//     "1764": {
//       "branch": null,
//       "build": 1764,
//       "files": [
//         [
//           "zip",
//           "mdk",
//           "7894ee158c6d95bd530c0fa9ca43fc9d"
//         ],
//         [
//           "txt",
//           "changelog",
//           "0e73961e23cc964f311a439969ce6b30"
//         ],
//         [
//           "jar",
//           "universal",
//           "416a1508f35e132cd36e199051757d14"
//         ],
//         [
//           "jar",
//           "userdev",
//           "68cfda844e19a519ecf51068b8450df6"
//         ],
//         [
//           "exe",
//           "installer-win",
//           "70e5d4fe8af98fc46da4a3479e1130bf"
//         ],
//         [
//           "jar",
//           "installer",
//           "a8a1be84b5a0b1ee90ef9e6f19d0c950"
//         ]
//       ],
//       "mcversion": "1.8.9",
//       "modified": 1456781231.0,
//       "version": "11.15.1.1764"
//     }
//   },
//   "promos": {

//   },
// }
