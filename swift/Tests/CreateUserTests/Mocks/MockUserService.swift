import Foundation
@testable import CreateUser

final class MockUserService: UserServiceProtocol {
    let result: Result<User, Error>
    init(result: Result<User, Error>) {
        self.result = result
    }
    func createUser(_ input: CreateUserInput) async throws -> User {
        switch result {
        case .success(let u):
            return u
        case .failure(let e):
            throw e
        }
    }
}
