import Foundation
@testable import CreateUser

final class MockRouter: CreateUserRouterProtocol {
    var lastUser: User?
    func showSuccess(from view: AnyObject, user: User) {
        lastUser = user
    }
}
