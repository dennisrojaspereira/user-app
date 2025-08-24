import XCTest
@testable import CreateUser

final class CreateUserTests: XCTestCase {
    func testPresenterValidatesEmptyName() {
        let service = MockUserService(result: .success(User(id: "1", name: "A", email: "a@b.com")))
        let interactor = CreateUserInteractor(service: service)
        let view = MockView()
        let router = MockRouter()
        let presenter = CreateUserPresenter(view: view, interactor: interactor, router: router)
        interactor.output = presenter
        presenter.didTapCreate(name: "", email: "a@b.com")
        XCTAssertEqual(view.messages.last, "Name is required.")
    }

    func testPresenterValidatesEmptyEmail() {
        let service = MockUserService(result: .success(User(id: "1", name: "A", email: "a@b.com")))
        let interactor = CreateUserInteractor(service: service)
        let view = MockView()
        let router = MockRouter()
        let presenter = CreateUserPresenter(view: view, interactor: interactor, router: router)
        interactor.output = presenter
        presenter.didTapCreate(name: "John", email: "")
        XCTAssertEqual(view.messages.last, "Email is required.")
    }

    func testInteractorSuccessCallsOutput() {
        let service = MockUserService(result: .success(User(id: "1", name: "John", email: "j@e.com")))
        let interactor = CreateUserInteractor(service: service)
        let exp = expectation(description: "didCreateUser called")
        final class Output: CreateUserInteractorOutput {
            let exp: XCTestExpectation
            init(_ exp: XCTestExpectation) { self.exp = exp }
            func didCreateUser(_ user: User) { exp.fulfill() }
            func didFailToCreateUser(_ error: Error) { XCTFail("Unexpected error") }
        }
        interactor.output = Output(exp)
        interactor.createUser(name: "John", email: "j@e.com")
        wait(for: [exp], timeout: 2.0)
    }

    func testPresenterSuccessFlow() {
        let user = User(id: "1", name: "John", email: "j@e.com")
        let service = MockUserService(result: .success(user))
        let interactor = CreateUserInteractor(service: service)
        let view = MockView()
        let router = MockRouter()
        let presenter = CreateUserPresenter(view: view, interactor: interactor, router: router)
        interactor.output = presenter
        presenter.didTapCreate(name: "John", email: "j@e.com")
        let exp = expectation(description: "async")
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) { exp.fulfill() }
        wait(for: [exp], timeout: 2.0)
        XCTAssertTrue(view.cleared)
        XCTAssertEqual(router.lastUser?.id, "1")
    }

    func testPresenterFailureFlow() {
        let service = MockUserService(result: .failure(UserServiceError.invalidEmail))
        let interactor = CreateUserInteractor(service: service)
        let view = MockView()
        let router = MockRouter()
        let presenter = CreateUserPresenter(view: view, interactor: interactor, router: router)
        interactor.output = presenter
        presenter.didTapCreate(name: "John", email: "invalid")
        let exp = expectation(description: "async")
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) { exp.fulfill() }
        wait(for: [exp], timeout: 2.0)
        XCTAssertEqual(view.messages.last, "Invalid email.")
    }
}
