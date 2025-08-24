import Foundation

public final class CreateUserInteractor: CreateUserInteractorInput {
    private let service: UserServiceProtocol
    public weak var output: CreateUserInteractorOutput?

    public init(service: UserServiceProtocol) {
        self.service = service
    }

    public func createUser(name: String, email: String) {
        Task { [weak self] in
            do {
                let user = try await self?.service.createUser(CreateUserInput(name: name, email: email))
                if let user = user {
                    await MainActor.run { self?.output?.didCreateUser(user) }
                }
            } catch {
                await MainActor.run { self?.output?.didFailToCreateUser(error) }
            }
        }
    }
}
