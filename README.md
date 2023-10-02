# GoVec
A library providing vector operations for Go.

The library supports both 2D and 3D vectors in two formats:
- Floating Point Vectors: `V2F[T]` and `V3F[T]`
- Integer Vectors: `V2I[T]` and `V3I[T]`

## Installation

```bash
go get github.com/xkonti/govec
```

## Usage

### Vector types

The library provides the following vector types:
- `V2F[T]` - 2D floating point vector, where `T` is the type of the vector components (`float32` or `float64`)
- `V2I[T]` - 2D integer vector, where `T` is the type of the vector components (`int`, `int8`, `int16`, `int32` or `int64`)
- `V3F[T]` - 3D floating point vector, where `T` is the type of the vector components (`float32` or `float64`)
- `V3I[T]` - 3D integer vector, where `T` is the type of the vector components (`int`, `int8`, `int16`, `int32` or `int64`)

### Creating new vectors

Vectors can be created in various ways:
- Constructing vectors manually:
    ```go
    vector1 := govec.V2F{1.0, 2.0};
    vector2 := govec.V3F{1.0, 2.0, 3.0};
    vector3 := govec.V2I{1, 2};
    vector4 := govec.V3I{1, 2, 3};
    ```
- Using the [`New`](new.go) function, which takes the vector components as arguments:
    ```go
    vector1 := govec.NewV2F(1.0, 2.0);
    vector2 := govec.NewV3F(1.0, 2.0, 3.0);
    vector3 := govec.NewV2I(1, 2);
    vector4 := govec.NewV3I(1, 2, 3);
    ```
- Using the [`Fill`](new.go) function, which assigns the same value to all components of a vector:
    ```go
    vector1 := govec.FillV2F(1.0);
    vector2 := govec.FillV3F(1.0);
    vector3 := govec.FillV2I(1);
    vector4 := govec.FillV3I(1);
    ```
- Using the [`Zero`](new.go) function, which creates a vector with all components set to 0:
    ```go
    vector1 := govec.ZeroV2F();
    vector2 := govec.ZeroV3F();
    vector3 := govec.ZeroV2I();
    vector4 := govec.ZeroV3I();
    ```
- Using the [`One`](new.go) function, which creates a vector with all components set to 1:
    ```go
    vector1 := govec.OneV2F();
    vector2 := govec.OneV3F();
    vector3 := govec.OneV2I();
    vector4 := govec.OneV3I();
    ```
- Using [`...FromArray`](array.go)
- Using [`...FromSlice`](slice.go)
  
### Conversions

Vectors can be converted to other types using the `To...` functions:
```go
vector1 := govec.V3F{1.0, 2.0, 3.0}; // Creates a 3D floating point vector (float64)
vector2 := vector1.ToV3I16(); // converts to 3D vector with int16 components
```

2D vectors have the following conversion functions:
- `ToV2F64`, `ToV2F32`, `ToV2I64`, `ToV2I32`, `ToV2I16`, `ToV2I8`, `ToV2I`

3D vectors have the following conversion functions:
- `ToV3F64`, `ToV3F32`, `ToV3I64`, `ToV3I32`, `ToV3I16`, `ToV3I8`, `ToV3I`

You can remove a component from a vector using the [`Discard`](discard.go) function:
```go
v1 := govec.V3F{1.0, 2.0, 3.0}; // Creates a 3D floating point vector (float64)
v2 := vector1.DiscardY(); // Returns a 2D vector with the Y component removed (v1.X -> v2.X, v1.Z -> v2.Y)
```

You can add a component to a vector using the [`Extend`](extend.go) function:
```go
v1 := govec.V2F{1.0, 2.0}; // Creates a 2D floating point vector (float64)
v2 := v1.Extend(3.0); // Returns a 3D vector with the Z component added.
```

You can insert a component to a vector using the [`Insert`](insert.go) function:
```go
v1 := govec.V2F{1.0, 2.0}; // Creates a 2D floating point vector (float64)
v2 := v1.InsertY(3.0); // Returns a 3D vector with the new component: newY -> Y, oldY -> Z.
```

You can swap components in a vector using the [`Swizzle`](swizzle.go) function:
```go
v1 := govec.V3F{1.0, 2.0, 3.0}; // Creates a 3D floating point vector (float64)
v2 := v1.SwizzleZYX(); // Returns a 3D vector with the components swapped (v1.Z -> v2.X, v1.Y -> v2.Y, v1.X -> v2.Z)
v3 := v1.SwizzleYZX(); // Returns a 3D vector with the components swapped (v1.Y -> v2.X, v1.Z -> v2.Y, v1.X -> v2.Z)
```

### Common variants

Most vector operations are available in various forms:
- Default form - returns a new vector, usually takes other vector as an argument
- `InPlace` form - modifies the original vector, usually takes other vector as an argument
- `Comp` form - returns a new vector, usually takes raw vector components as arguments instead of another vector
- `CompInPlace` form - modifies the original vector, usually takes raw vector components as arguments instead of another vector
- `Scalar` form - returns a new vector, takes a scalar value as an argument
- `ScalarInPlace` form - modifies the original vector, takes a scalar value as an argument

For example, the [`Add`](add.go) operation is available as:
- [`Add`](add.go) - takes another vector as an argument and returns a new vector
- [`AddInPlace`](add.go) - takes another vector as an argument and modifies the original vector
- [`AddComp`](add.go) - takes raw vector components (x, y, z) as arguments and returns a new vector
- [`AddCompInPlace`](add.go) - takes raw vector components (x, y, z) as arguments and modifies the original vector
- [`AddScalar`](addScalar.go) - takes a scalar value as an argument and returns a new vector
- [`AddScalarInPlace`](addScalar.go) - takes a scalar value as an argument and modifies the original vector

For some operations, the `Comp`, `InPlace` or `Scalar` variants are not available as it doesn't make sense to use them.

### Return type exceptions

Usually the return type of the operation is the same as the type of the vector on which the operation is performed. However,
in some cases the return type is different. Examples:
- `Len` operation always returns `float64` even if the vector is an integer vector. This is the only way to return an accurate
  length of the vector.
- `Norm` operations always returns a `float64` vectors even if the original vector is an integer vector. Normalization
  returns values between 0 and 1, which is not possible to represent in an integer vector.

## Available Operations

| Operation                   | Float | Integer | Description                                                                                                 |
|-----------------------------|-------|---------|-------------------------------------------------------------------------------------------------------------|
| **Creation**                |       |         |                                                                                                             |
| [`...FromArray`](array.go)  | ✔️    | ✔️      | Initializes a vector from an array of components.                                                           |
| [`...FromSlice`](slice.go)  | ✔️    | ✔️      | Initializes a vector from a slice of components.                                                            |
| [`Fill`](new.go)            | ✔️    | ✔️      | Creates a vector where all components are set to a specified value.                                         |
| [`New`](new.go)             | ✔️    | ✔️      | Creates a new vector using the provided components.                                                         |
| [`One`](new.go)             | ✔️    | ✔️      | Generates a vector with all components set to 1.                                                            |
| [`Zero`](new.go)            | ✔️    | ✔️      | Generates a vector with all components set to 0.                                                            |
| **Conversions**             |       |         |                                                                                                             |
| [`Apply`](apply.go)         | ✔️    | ✔️      | Applies a given function to each component in the vector.                                                   |
| [`ApplyToArray`](array.go)  | ✔️    | ✔️      | Modifies an array by assigning values from the vector's corresponding components.                           |
| [`ApplyToSlice`](slice.go)  | ✔️    | ✔️      | Modifies a slice by assigning values from the vector's corresponding components.                            |
| [`Discard`](discard.go)     | ✔️    | ✔️      | Removes a component from the vector at a specified index.                                                   |
| [`Extend`](extend.go)       | ✔️    | ✔️      | Extends the dimensionality of a vector by adding additional components.                                     |
| [`Insert`](insert.go)       | ✔️    | ✔️      | Inserts a value into the vector at a specified index.                                                       |
| [`Split`](split.go)         | ✔️    | ✔️      | Splits a vector into its individual components.                                                             |
| [`Swizzle...`](swizzle.go)  | ✔️    | ✔️      | Swaps components of a vector using a selected pattern.                                                      |
| [`ToArray`](array.go)       | ✔️    | ✔️      | Converts the vector to a new array.                                                                         |
| [`ToSlice`](slice.go)       | ✔️    | ✔️      | Converts the vector to a new slice.                                                                         |
| `To...`                     | ✔️    | ✔️      | Transforms the vector into another type. [V2F](vec2f.go), [V3F](vec3f.go), [V2I](vec2i.go), [V3I](vec3i.go) |
| **Mathematics**             |       |         |                                                                                                             |
| [`Abs`](abs.go)             | ✔️    | ✔️      | Computes the absolute value of each component in the vector.                                                |
| [`Add`](add.go)             | ✔️    | ✔️      | Performs vector addition.                                                                                   |
| [`AddScalar`](addScalar.go) | ✔️    | ✔️      | Adds a scalar value to each component of the vector.                                                        |
| [`Ceil`](ceil.go)           | ✔️    | ✔️      | Rounds each component of the vector up to the nearest integer.                                              |
| [`ClampComp`](clampComp.go) | ✔️    | ✔️      | Clamps the components of this vector to the given range.                                                    |
| [`ClampLen`](clampLen.go)   | ✔️    | ✔️      | Scales the length of the vector so that it fits within provided range.                                      |
| [`Cross`](cross.go)         | ✔️    | ✔️      | Computes the cross product of two vectors. (Not applicable for 2D vectors.)                                 |
| [`Div`](div.go)             | ✔️    | ✔️      | Performs vector division.                                                                                   |
| [`DivScalar`](dicScalar.go) | ✔️    | ✔️      | Divides each component of the vector by a scalar.                                                           |
| [`Dot`](dot.go)             | ✔️    | ✔️      | Computes the dot product of two vectors.                                                                    |
| [`Floor`](floor.go)         | ✔️    | ✔️      | Rounds each component of the vector down to the nearest integer.                                            |
| [`Inv`](inv.go)             | ✔️    | ✔️      | Computes the multiplicative inverse of each component in the vector.                                        |
| [`Len`](len.go)             | ✔️    | ✔️      | Calculates the length of the vector. Returns `float64` for integer vectors.                                 |
| [`LenSqrt`](lenSqrt.go)     | ✔️    | ✔️      | Calculates the squared length of the vector. Returns `float64` for integer vectors.                         |
| [`Max`](max.go)             | ✔️    | ✔️      | Returns the maximum component values from two vectors.                                                      |
| [`Min`](min.go)             | ✔️    | ✔️      | Returns the minimum component values from two vectors.                                                      |
| [`Mod`](mod.go)             | ✔️    | ✔️      | Computes the modulus of each component in the vector against another vector components.                     |
| [`ModScalar`](modScalar.go) | ✔️    | ✔️      | Computes the modulus of each component in the vector against a single scalar.                               |
| [`Mul`](mul.go)             | ✔️    | ✔️      | Performs vector multiplication.                                                                             |
| [`MulScalar`](mulScalar.go) | ✔️    | ✔️      | Multiplies each component of the vector by a scalar.                                                        |
| [`Neg`](neg.go)             | ✔️    | ✔️      | Negates each component of the vector.                                                                       |
| [`Norm`](norm.go)           | ✔️    | ✔️      | Normalizes the vector. Returns a `float64` vector for integer vectors.                                      |
| [`Pow`](pow.go)             | ✔️    | ✔️      | Raises each component of the vector to the power of the corresponding component in another vector.          |
| [`Pow2`](pow2.go)           | ✔️    | ✔️      | Squares each component of the vector.                                                                       |
| [`Pow3`](pow3.go)           | ✔️    | ✔️      | Cubes each component of the vector.                                                                         |
| [`PowN`](powN.go)           | ✔️    | ✔️      | Raises each component of the vector to a specified integer power.                                           |
| [`PowNFloat`](powNFloat.go) | ✔️    | ✔️      | Raises each component of the vector to a specified floating-point power.                                    |
| [`Round`](round.go)         | ✔️    | ✔️      | Rounds each component of the vector to the nearest integer.                                                 |
| [`Sqrt`](sqrt.go)           | ✔️    | ✔️      | Computes the square root of each component in the vector.                                                   |
| [`Sub`](sub.go)             | ✔️    | ✔️      | Performs vector subtraction.                                                                                |
| [`SubScalar`](subScalar.go) | ✔️    | ✔️      | Subtracts a scalar from each component of the vector.                                                       |

## Not implemented yet

The following list of operations are not implemented yet. Feel free to open an issue if you would like to see
any of these implemented or if you'd like to propose a different operation.

| Operation         | Planned    | Description                                                                       |
|-------------------|------------|-----------------------------------------------------------------------------------|
| `AngleBetweenDeg` | Yes (v1.0) | Calculates the angle between two vectors in degrees.                              |
| `AngleBetweenRad` | Yes (v1.0) | Calculates the angle between two vectors in radians.                              |
| `AngleDeg`        | Yes (v1.0) | Creates a normalized vector from an angle in degrees.                             |
| `AngleRad`        | Yes (v1.0) | Creates a normalized vector from an angle in radians.                             |
| `Average`         | Yes (v1.0) | Calculates the average of the vector's components.                                |
| `Cos`             |            | Applies the cosine function to all components.                                    |
| `Distance`        | Yes (v1.0) | Calculates the distance between two vectors.                                      |
| `FromQuaternion`  |            | Initializes the vector from a quaternion.                                         |
| `Hash`            |            | Produces a hash code for the vector for use in data structures.                   |
| `IsOrthogonalTo`  |            | Checks if the vector is orthogonal to another vector.                             |
| `IsUnit`          |            | Checks if the vector is a unit vector.                                            |
| `IsZero`          | Yes (v1.0) | Checks if the vector is a zero vector.                                            |
| `Lerp`            | Yes (v1.0) | Linearly interpolates between two vectors.                                        |
| `Orthogonalize`   |            | Generates an orthogonal (or orthonormal) vector set.                              |
| `Project`         |            | Projects a 3D vector onto a plane.                                                |
| `Rand`            | Yes (v1.0) | Generates a normal vector with random components.                                 |
| `RandIn`          | Yes (v1.0) | Generates a vector with random components within a zero and the specified vector. |
| `RandBetween`     | Yes (v1.0) | Generates a vector with random components between two specified vectors.          |
| `Reflect`         |            | Reflects a vector off the plane defined by a normal.                              |
| `RotateDeg`       | Yes (v1.0) | Rotates a vector by an angle in degrees.                                          |
| `RotateRad`       | Yes (v1.0) | Rotates a vector by an angle in radians.                                          |
| `Sin`             |            | Applies the sine function to all components.                                      |
| `Slerp`           |            | Spherically interpolates between two vectors.                                     |
| `Tan`             |            | Applies the tangent function to all components.                                   |
| `ToQuaternion`    |            | Converts the vector to a quaternion.                                              |