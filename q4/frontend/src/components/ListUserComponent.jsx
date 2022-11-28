import React, { Component } from 'react'
import { DataGrid } from '@mui/x-data-grid';
import UserService from '../services/UserService'

const columns = [
    { field: 'firstName', headerName: 'First Name', width: 150 },
    { field: 'lastName', headerName: 'Last Name', width: 150 },
    { field: 'email', headerName: 'email', width: 150 },
];

class ListUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            users: [],
            selectedUserId: '',
            type: this.props.match.params.type,
            id: this.props.match.params.id
        }
        this.addUser = this.addUser.bind(this);
        this.editUser = this.editUser.bind(this);
        this.deleteUser = this.deleteUser.bind(this);

        // If the user list is displayed after deleting a user, do not request all users from database again.
        // Just delete the deleted used from this.state.users
        if (this.state.type === 'delete') {
            this.setState({users: this.state.users.filter(user => user.id !== this.state.id)});
        }
    }

    addUser(){
        this.props.history.push(`/action/add/0`);
    }

    editUser(){
        this.props.history.push(`/action/edit/${this.state.selectedUserId}`);
    }

    deleteUser(){
        this.props.history.push(`/action/delete/${this.state.selectedUserId}`);
    }

    componentDidMount(){
        UserService.getUsers().then((res) => {
            this.setState({users: res.data || []});
        });
    }

    render() {
        return (
            <div>
                <h2 className="text-center">
                    Users List</h2>
                <div className = "row">
                    <button className="btn btn-primary"
                        onClick={this.addUser}
                        style={{ marginLeft: "15px" }}
                    >
                        New
                    </button>
                    <button className="btn btn-info"
                        onClick={this.editUser}
                        style={{ marginLeft: "15px" }}
                        disabled={!this.state.selectedUserId}
                    >
                        Edit
                    </button>
                    <button className="btn btn-danger"
                        onClick={this.deleteUser}
                        style={{ marginLeft: "15px" }}
                        disabled={!this.state.selectedUserId}
                    >
                        Delete
                    </button>
                </div>
                <br></br>
                <div style={{ height: 300, width: '100%' }}>
                    <DataGrid
                        checkboxSelection={true}
                        isRowSelectable={(params) =>
                            this.state.selectedUserId === '' || params.id === this.state.selectedUserId
                        }
                        onSelectionModelChange={(selectedUserIds) => {
                            if (selectedUserIds && selectedUserIds.length) {
                                this.setState({selectedUserId: selectedUserIds[0]})
                            } else {
                                this.setState({selectedUserId: ''})
                            }
                        }}
                        rows={this.state.users} columns={columns}
                    />
                </div>
            </div>
        )
    }
}

export default ListUserComponent
