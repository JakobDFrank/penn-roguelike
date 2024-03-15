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

4. Run `npm start` to start the React server.

5. In another terminal, run `npm run electron-dev` to start the GUI.

**To do: Make production ready and improve documentation. This is developer only at the moment.**

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