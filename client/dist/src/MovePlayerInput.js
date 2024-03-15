import { ADDRESS, PLAYER_MOVE_ENDPOINT } from "./App";
export var Direction;
(function (Direction) {
    Direction[Direction["Left"] = 0] = "Left";
    Direction[Direction["Right"] = 1] = "Right";
    Direction[Direction["Up"] = 2] = "Up";
    Direction[Direction["Down"] = 3] = "Down";
})(Direction || (Direction = {}));
export function MovePlayerInput(setNums, id, dir) {
    var num;
    switch (dir) {
        case Direction.Left:
            num = 0;
            break;
        case Direction.Right:
            num = 1;
            break;
        case Direction.Up:
            num = 2;
            break;
        case Direction.Down:
            num = 3;
            break;
        default:
            return null;
    }
    var req = {
        id: id,
        direction: num,
    };
    console.log(req);
    var url = ADDRESS + PLAYER_MOVE_ENDPOINT;
    fetch(url, {
        method: "POST",
        body: JSON.stringify(req),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    })
        .then(function (response) { return response.json(); })
        .then(function (data) {
        console.log(data);
        if (data.level) {
            setNums(data.level);
        }
        else {
            console.error("newBoard error");
        }
    })
        .catch(function (err) {
        console.log(err.message);
    });
}
