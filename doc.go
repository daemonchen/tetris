// package tetris implements a well-known game tetris
//
// Overview
//
// The coordinate system
// The width and height of the game screen is not const value, you can set the values as you like.
// The origin is located on the top-left and the x-axis grows from left to right while y-axis grows from top to bottom.
//
// The inherit relation between the data types in the package:
// (well, maybe it is not correct to say inherit, containment could be better)
// 1. gameController -> gameZone, timer, combos
// 2. gameZone -> piece, line
// 3. piece -> dots -> dot
//
// The dot type represents a coordinate(x, y) on a 2D space
// It is able to rotate a dot by another dot for the convinience purpose of piece roration
//
// The dots type represents a piece in a relative coordinate
// It is a collection of dot and provide some help functions for piece type.
//
// The piece type represents a piece in a absolute coordinate
//
// More detail... Too tired of documentation
package tetris
