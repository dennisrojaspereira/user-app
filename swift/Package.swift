// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "CreateUserVIPER",
    platforms: [.iOS(.v15)],
    products: [
        .library(name: "CreateUser", targets: ["CreateUser"])
    ],
    targets: [
        .target(
            name: "CreateUser",
            path: "Sources/CreateUser"
        ),
        .testTarget(
            name: "CreateUserTests",
            dependencies: ["CreateUser"],
            path: "Tests/CreateUserTests"
        )
    ]
)
