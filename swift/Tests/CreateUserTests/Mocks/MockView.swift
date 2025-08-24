import Foundation
@testable import CreateUser

final class MockView: CreateUserViewProtocol {
    var isLoading: Bool = false
    var messages: [String] = []
    var cleared: Bool = false
    func setLoading(_ loading: Bool) { isLoading = loading }
    func showMessage(_ text: String) { messages.append(text) }
    func clearFields() { cleared = true }
}
