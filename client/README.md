# penn-roguelike GUI client

---

## Prerequisites

- React

- Typescript

- Electron

## How to Use

1. Clone the repository.

2. Follow the instructions at [penn-roguelike/server](../server/README.md) and run an HTTP server.

3. Open a terminal and navigate to penn-roguelike/client.

4. To run the application:
   - For Development: Run `npm run electron-dev` to start the application in development mode.
   - For Compilation: Run `npm run electron:package:linux` to build the app for Linux x86_64. <br>Find the AppImage at penn-roguelike/client/dist/client-0.1.0.AppImage.

### Submitting a Level

You must submit a level through "Load" to play the game.

Any submitted level will immediately be used as the active level.

- `0` for an open tile
- `1` for a wall
- `2` for a pit (player takes one damage)
- `3` for arrows (player takes two damage)
- `4` for the player

Constraints:
- Map must be rectangular
- Map height and map width can not exceed 100 units
- Map spaces can only consist of zero to four, inclusive
- There is one, and only one, player

### Moving the Player

You can move the player using W,A,S,D or the arrow keys.

The level must be clicked first.