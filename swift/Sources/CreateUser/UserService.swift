import Foundation

public protocol UserServiceProtocol {
    func createUser(_ input: CreateUserInput) async throws -> User
}

public enum UserServiceError: Error, LocalizedError, Equatable {
    case invalidEmail
    case server(String)
    public var errorDescription: String? {
        switch self {
        case .invalidEmail: return "Invalid email."
        case .server(let msg): return msg
        }
    }
}

public final class UserService: UserServiceProtocol {
    public init() {}
    public func createUser(_ input: CreateUserInput) async throws -> User {
        guard input.email.contains("@") else { throw UserServiceError.invalidEmail }
        try await Task.sleep(nanoseconds: 200_000_000)
        return User(id: UUID().uuidString, name: input.name, email: input.email)
    }
}
