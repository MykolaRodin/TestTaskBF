import React, { Component } from 'react'
import UserService from '../services/UserService';

class CreateUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            type: this.props.match.params.type,
            id: this.props.match.params.id,
            firstName: '',
            lastName: '',
            email: ''
        }
        this.changeFirstNameHandler =
            this.changeFirstNameHandler.bind(this);
        this.changeLastNameHandler =
            this.changeLastNameHandler.bind(this);
        this.saveOrUpdateUser =
            this.saveOrUpdateUser.bind(this);
    }

    // step 3
    componentDidMount() {

        // step 4
        if (this.state.type === 'add') {
            return
        } else { // For both 'edit' and 'delete'
            UserService.getUserById(this.state.id).
            then((res) => {
                let user = res.data;
                this.setState({
                    firstName: user.firstName,
                    lastName: user.lastName,
                    email: user.email
                });
            });
        }
    }
    saveOrUpdateUser = (e) => {
        e.preventDefault();
        let user = { firstName: this.state.firstName, lastName:
             this.state.lastName, email: this.state.email };
        console.log('user => ' + JSON.stringify(user));

        // step 5
        if (this.state.type === 'add') {
            UserService.createUser(user).then(() => {
                this.props.history.push(`/users/none/0`);
            });
        } else if (this.state.type === 'delete') {
            UserService.deleteUser(this.state.id).then(() => {
                this.props.history.push(`/users/delete/${this.state.id}`);
            });
        } else { // 'edit'
            UserService.updateUser(user, this.state.id).then(() => {
                this.props.history.push(`/users/edit/${this.state.id}`);
            });
        }
    }

    changeFirstNameHandler = (event) => {
        this.setState({ firstName: event.target.value });
    }

    changeLastNameHandler = (event) => {
        this.setState({ lastName: event.target.value });
    }

    changeEmailHandler = (event) => {
        this.setState({ email: event.target.value });
    }

    cancel() {
        this.props.history.push('/users/none/0');
    }

    getTitle() {
        const labelText = this.state.type === 'add' ? 'Add User'
                        : this.state.type === 'delete' ? 'Delete User'
                        : 'Edit User';
        return <h3 className="text-center">{labelText}</h3>
    }
    render() {
        const disabledEdit = this.state.type === 'delete';
        const mainButtonText = this.state.type === 'add' ? 'Create'
                            : this.state.type === 'delete' ? 'Delete'
                            : 'Save';
        return (
            <div>
                <br></br>
                <div className="container">
                    <div className="row">
                        <div className="card col-md-6 offset-md-3 offset-md-3">
                            {
                                this.getTitle()
                            }
                            <div className="card-body">
                                <form>
                                    <div className="form-group">
                                        <label> First Name: </label>
                                        <input placeholder="First Name"
                                            name="firstName" className="form-control"
                                            value={this.state.firstName}
                                            onChange={this.changeFirstNameHandler}
                                            disabled={disabledEdit}
                                        />
                                    </div>
                                    <div className="form-group">
                                        <label> Last Name: </label>
                                        <input placeholder="Last Name"
                                            name="lastName" className="form-control"
                                            value={this.state.lastName}
                                            onChange={this.changeLastNameHandler}
                                            disabled={disabledEdit}
                                        />
                                    </div>
                                    <div className="form-group">
                                        <label> Email : </label>
                                        <input placeholder="Email Address"
                                            name="email" className="form-control"
                                            value={this.state.email}
                                            onChange={this.changeEmailHandler}
                                            disabled={disabledEdit}
                                        />
                                    </div>
                                    <button className="btn btn-success"
                                        onClick={this.saveOrUpdateUser}
                                    >
                                        {mainButtonText}
                                    </button>
                                    <button className="btn btn-danger"
                                        onClick={this.cancel.bind(this)}
                                        style={{ marginLeft: "10px" }}
                                    >
                                        Back
                                    </button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default CreateUserComponent
