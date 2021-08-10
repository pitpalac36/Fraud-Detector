import React from 'react';
import 'antd/dist/antd.css';
import './index.css';
import { Pagination } from 'antd';

export class MyPagination extends React.Component {
    state = {
        current: 3,
    };

    onChange = (page: any) => {
        console.log(page);
        this.setState({
            current: page,
        });
    };

    render() {
        return <Pagination current={this.state.current} onChange={this.onChange} defaultPageSize={20}/>;
    }
}