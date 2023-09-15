# GoVec
A library providing vector operations for Go.

The library supports both 2D and 3D vectors in two formats:
- Floating Point Vectors: `V2F[T]` and `V3F[T]`
- Integer Vectors: `V2I[T]` and `V3I[T]`

## Installation

```bash
go get github.com/xkonti/govec
```

## Available Operations

| Operation       | Float | Integer | Description                                                                                        |
|-----------------|----|----|----------------------------------------------------------------------------------------------------|
| **Creation**    |    |    |                                                                                                    |
| `...FromArray`  | ✔️ | ✔️ | Initializes a vector from an array of components.                                                  |
| `...FromSlice`  | ✔️ | ✔️ | Initializes a vector from a slice of components.                                                   |
| `Fill`          | ✔️ | ✔️ | Creates a vector where all components are set to a specified value.                                |
| `New`           | ✔️ | ✔️ | Creates a new vector using the provided components.                                                |
| `One`           | ✔️ | ✔️ | Generates a vector with all components set to 1.                                                   |
| `Zero`          | ✔️ | ✔️ | Generates a vector with all components set to 0.                                                   |
| **Conversions** |    |    |                                                                                                    |
| `Apply`         | ✔️ | ✔️ | Applies a given function to each component in the vector.                                          |
| `ApplyToArray`  | ✔️ | ✔️ | Modifies an array by assigning values from the vector's corresponding components.                  |
| `ApplyToSlice`  | ✔️ | ✔️ | Modifies a slice by assigning values from the vector's corresponding components.                   |
| `Discard`       | ✔️ | ✔️ | Removes a component from the vector at a specified index.                                          |
| `Extend`        | ✔️ | ✔️ | Extends the dimensionality of a vector by adding additional components.                            |
| `Insert`        | ✔️ | ✔️ | Inserts a value into the vector at a specified index.                                              |
| `Split`         | ✔️ | ✔️ | Splits a vector into its individual components.                                                    |
| `ToArray`       | ✔️ | ✔️ | Converts the vector to a new array.                                                                |
| `ToSlice`       | ✔️ | ✔️ | Converts the vector to a new slice.                                                                |
| `To...`         | ✔️ | ✔️ | Transforms the vector into another type.                                                           |
| **Mathematics** |    |    |                                                                                                    |
| `Abs`           | ✔️ | ✔️ | Computes the absolute value of each component in the vector.                                       |
| `Add`           | ✔️ | ✔️ | Performs vector addition.                                                                          |
| `AddScalar`     | ✔️ | ✔️ | Adds a scalar value to each component of the vector.                                               |
| `Ceil`          | ✔️ | ✔️ | Rounds each component of the vector up to the nearest integer.                                     |
| `ClampComp`     | ✔️ | ✔️ | Clamps the components of this vector to the given range.                                           |
| `ClampLen`      | ✔️ | ✔️ | Scales the length of the vector so that it fits within provided range.                             |
| `Cross`         | ✔️ | ✔️ | Computes the cross product of two vectors. (Not applicable for 2D vectors.)                        |
| `Div`           | ✔️ | ✔️ | Performs vector division.                                                                          |
| `DivScalar`     | ✔️ | ✔️ | Divides each component of the vector by a scalar.                                                  |
| `Dot`           | ✔️ | ✔️ | Computes the dot product of two vectors.                                                           |
| `Floor`         | ✔️ | ✔️ | Rounds each component of the vector down to the nearest integer.                                   |
| `Inv`           | ✔️ | ✔️ | Computes the multiplicative inverse of each component in the vector.                               |
| `Len`           | ✔️ | ✔️ | Calculates the length of the vector. Returns `float64` for integer vectors.                        |
| `LenSqr`        | ✔️ | ✔️ | Calculates the squared length of the vector. Returns `float64` for integer vectors.                |
| `Max`           | ✔️ | ✔️ | Returns the maximum component values from two vectors.                                             |
| `Min`           | ✔️ | ✔️ | Returns the minimum component values from two vectors.                                             |
| `Mod`           | ✔️ | ✔️ | Computes the modulus of each component in the vector against another vector components.            |
| `ModScalar`     | ✔️ | ✔️ | Computes the modulus of each component in the vector against a single scalar.                      |
| `Mul`           | ✔️ | ✔️ | Performs vector multiplication.                                                                    |
| `MulScalar`     | ✔️ | ✔️ | Multiplies each component of the vector by a scalar.                                               |
| `Neg`           | ✔️ | ✔️ | Negates each component of the vector.                                                              |
| `Norm`          | ✔️ | ✔️ | Normalizes the vector. Returns a `float64` vector for integer vectors.                             |
| `Pow`           | ✔️ | ✔️ | Raises each component of the vector to the power of the corresponding component in another vector. |
| `Pow2`          | ✔️ | ✔️ | Squares each component of the vector.                                                              |
| `Pow3`          | ✔️ | ✔️ | Cubes each component of the vector.                                                                |
| `PowN`          | ✔️ | ✔️ | Raises each component of the vector to a specified integer power.                                  |
| `PowNFloat`     | ✔️ | ✔️ | Raises each component of the vector to a specified floating-point power.                           |
| `Round`         | ✔️ | ✔️ | Rounds each component of the vector to the nearest integer.                                        |
| `Sqrt`          | ✔️ | ✔️ | Computes the square root of each component in the vector.                                          |
| `Sub`           | ✔️ | ✔️ | Performs vector subtraction.                                                                       |
| `SubScalar`     | ✔️ | ✔️ | Subtracts a scalar from each component of the vector.                                              |

## Not implemented yet

The following list of operations are not implemented yet. Feel free to open an issue if you would like to see
any of these implemented or if you'd like to propose a different operation.

| Operation         | Planned    | Description                                                 |
|-------------------|------------|-------------------------------------------------------------|
| `AngleBetweenDeg` | Yes (v1.0) | Calculates the angle between two vectors in degrees.        |
| `AngleBetweenRad` | Yes (v1.0) | Calculates the angle between two vectors in radians.        |
| `AngleDeg`        | Yes (v1.0) | Creates a normalized vector from an angle in degrees.       |
| `AngleRad`        | Yes (v1.0) | Creates a normalized vector from an angle in radians.       |
| `Average`         | Yes (v1.0) | Calculates the average of the vector's components.          |
| `Cos`             |            | Applies the cosine function to all components.              |
| `Distance`        | Yes (v1.0) | Calculates the distance between two vectors.                |
| `FromQuaternion`  |            | Initializes the vector from a quaternion.                   |
| `Hash`            |            | Produces a hash code for the vector for use in data structures. |
| `IsOrthogonalTo`  |            | Checks if the vector is orthogonal to another vector.       |
| `IsUnit`          |            | Checks if the vector is a unit vector.                      |
| `IsZero`          | Yes (v1.0) | Checks if the vector is a zero vector.                      |
| `Lerp`            | Yes (v1.0) | Linearly interpolates between two vectors.                  |
| `Orthogonalize`   |            | Generates an orthogonal (or orthonormal) vector set.        |
| `Project`         |            | Projects a 3D vector onto a plane.                          |
| `Random`          | Yes (v1.0) | Generates a normal vector with random components.           |
| `Reflect`         |            | Reflects a vector off the plane defined by a normal.        |
| `RotateDeg`       | Yes (v1.0) | Rotates a vector by an angle in degrees.                    |
| `RotateRad`       | Yes (v1.0) | Rotates a vector by an angle in radians.                    |
| `Sin`             |            | Applies the sine function to all components.                |
| `Slerp`           |            | Spherically interpolates between two vectors.               |
| `Tan`             |            | Applies the tangent function to all components.             |
| `ToQuaternion`    |            | Converts the vector to a quaternion.                        |