import Foundation
#if canImport(UIKit)
import UIKit
#endif

public enum CreateUserModule {
    #if canImport(UIKit)
    public static func build(service: UserServiceProtocol = UserService()) -> UIViewController {
        let view = CreateUserViewController()
        let interactor = CreateUserInteractor(service: service)
        let router = CreateUserRouter()
        let presenter = CreateUserPresenter(view: view, interactor: interactor, router: router)
        interactor.output = presenter
        view.presenter = presenter
        return UINavigationController(rootViewController: view)
    }
    #endif
}
