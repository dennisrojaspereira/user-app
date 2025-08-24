import Foundation
#if canImport(UIKit)
import UIKit
#endif

public protocol CreateUserViewProtocol: AnyObject {
    func setLoading(_ loading: Bool)
    func showMessage(_ text: String)
    func clearFields()
}

public protocol CreateUserPresenterProtocol: AnyObject {
    func didTapCreate(name: String?, email: String?)
}

public protocol CreateUserInteractorInput: AnyObject {
    func createUser(name: String, email: String)
}

public protocol CreateUserInteractorOutput: AnyObject {
    func didCreateUser(_ user: User)
    func didFailToCreateUser(_ error: Error)
}

public protocol CreateUserRouterProtocol: AnyObject {
    #if canImport(UIKit)
    func showSuccess(from view: UIViewController, user: User)
    #endif
}
