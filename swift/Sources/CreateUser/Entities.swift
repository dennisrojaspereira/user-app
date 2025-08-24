import Foundation

public struct User: Codable, Equatable {
    public let id: String
    public let name: String
    public let email: String
    public init(id: String, name: String, email: String) {
        self.id = id
        self.name = name
        self.email = email
    }
}

public struct CreateUserInput: Equatable {
    public let name: String
    public let email: String
    public init(name: String, email: String) {
        self.name = name
        self.email = email
    }
}
