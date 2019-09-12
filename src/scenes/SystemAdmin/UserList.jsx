import React from 'react';
import { List, Datagrid, TextField, BooleanField } from 'react-admin';
import AdminPagination from './AdminPagination';

const UserList = props => (
  <List {...props} pagination={<AdminPagination />} perPage={25}>
    <Datagrid rowClick="show">
      <TextField source="id" />
      <TextField source="email" />
      <TextField source="first_name" />
      <TextField source="last_name" />
      <BooleanField source="disabled" label="Deactivated" />
    </Datagrid>
  </List>
);

export default UserList;
