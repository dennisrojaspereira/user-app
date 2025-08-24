import Foundation
#if canImport(UIKit)
import UIKit
#endif

public final class CreateUserPresenter: CreateUserPresenterProtocol, CreateUserInteractorOutput {
    public weak var view: CreateUserViewProtocol?
    private let interactor: CreateUserInteractorInput
    private let router: CreateUserRouterProtocol

    public init(view: CreateUserViewProtocol, interactor: CreateUserInteractorInput, router: CreateUserRouterProtocol) {
        self.view = view
        self.interactor = interactor
        self.router = router
    }

    public func didTapCreate(name: String?, email: String?) {
        let trimmedName = (name ?? "").trimmingCharacters(in: .whitespacesAndNewlines)
        let trimmedEmail = (email ?? "").trimmingCharacters(in: .whitespacesAndNewlines)
        guard !trimmedName.isEmpty else { view?.showMessage("Name is required."); return }
        guard !trimmedEmail.isEmpty else { view?.showMessage("Email is required."); return }
        view?.setLoading(true)
        interactor.createUser(name: trimmedName, email: trimmedEmail)
    }

    public func didCreateUser(_ user: User) {
        view?.setLoading(false)
        view?.clearFields()
        #if canImport(UIKit)
        if let vc = view as? UIViewController {
            router.showSuccess(from: vc, user: user)
            return
        }
        #endif
        view?.showMessage("User created: \(user.name)")
    }

    public func didFailToCreateUser(_ error: Error) {
        view?.setLoading(false)
        view?.showMessage(error.localizedDescription)
    }
}
