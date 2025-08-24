#if canImport(UIKit)
import UIKit

public final class CreateUserRouter: CreateUserRouterProtocol {
    public init() {}
    public func showSuccess(from view: UIViewController, user: User) {
        let alert = UIAlertController(title: "Success", message: "User \(user.name) created.\nID: \(user.id)", preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .default))
        view.present(alert, animated: true)
    }
}
#endif
