#if canImport(UIKit)
import UIKit

public final class CreateUserViewController: UIViewController, CreateUserViewProtocol {
    private let nameField: UITextField = {
        let tf = UITextField()
        tf.placeholder = "Name"
        tf.borderStyle = .roundedRect
        tf.autocapitalizationType = .words
        return tf
    }()
    private let emailField: UITextField = {
        let tf = UITextField()
        tf.placeholder = "Email"
        tf.borderStyle = .roundedRect
        tf.keyboardType = .emailAddress
        tf.autocapitalizationType = .none
        return tf
    }()
    private let createButton: UIButton = {
        let bt = UIButton(type: .system)
        bt.setTitle("Create User", for: .normal)
        return bt
    }()
    private let spinner = UIActivityIndicatorView(style: .medium)

    public var presenter: CreateUserPresenterProtocol!

    public override func viewDidLoad() {
        super.viewDidLoad()
        title = "New User"
        view.backgroundColor = .systemBackground
        layout()
        createButton.addTarget(self, action: #selector(didTapCreate), for: .touchUpInside)
    }

    private func layout() {
        [nameField, emailField, createButton, spinner].forEach {
            $0.translatesAutoresizingMaskIntoConstraints = false
            view.addSubview($0)
        }
        NSLayoutConstraint.activate([
            nameField.topAnchor.constraint(equalTo: view.safeAreaLayoutGuide.topAnchor, constant: 24),
            nameField.leadingAnchor.constraint(equalTo: view.leadingAnchor, constant: 16),
            nameField.trailingAnchor.constraint(equalTo: view.trailingAnchor, constant: -16),
            emailField.topAnchor.constraint(equalTo: nameField.bottomAnchor, constant: 12),
            emailField.leadingAnchor.constraint(equalTo: nameField.leadingAnchor),
            emailField.trailingAnchor.constraint(equalTo: nameField.trailingAnchor),
            createButton.topAnchor.constraint(equalTo: emailField.bottomAnchor, constant: 20),
            createButton.centerXAnchor.constraint(equalTo: view.centerXAnchor),
            spinner.topAnchor.constraint(equalTo: createButton.bottomAnchor, constant: 16),
            spinner.centerXAnchor.constraint(equalTo: view.centerXAnchor)
        ])
    }

    @objc private func didTapCreate() {
        presenter.didTapCreate(name: nameField.text, email: emailField.text)
    }

    public func setLoading(_ loading: Bool) {
        if loading { spinner.startAnimating() } else { spinner.stopAnimating() }
        view.isUserInteractionEnabled = !loading
    }

    public func showMessage(_ text: String) {
        let alert = UIAlertController(title: nil, message: text, preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .default))
        present(alert, animated: true)
    }

    public func clearFields() {
        nameField.text = nil
        emailField.text = nil
    }
}
#endif
