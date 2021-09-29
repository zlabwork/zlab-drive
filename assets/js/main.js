let utilsBox = function () {

    let httpRequest = function (method, url, data) {
        // TODO :: configs
        axios.defaults.baseURL = 'http://localhost:8000/';
        axios({
            method: method,
            url: url,
            headers: {'content-type': 'application/x-www-form-urlencoded'},
            data: data
        })
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
            .then(function () {
                // always executed
            });
    }

    return {
        httpRequest: function (method, uri, data) {
            return httpRequest(method, uri, data)
        }
    }

}()

class Breadcrumb extends React.Component {
    constructor(props) {
        super(props);
        this.state = {date: new Date()};

        // 这边绑定是必要的，这样 `this` 才能在回调函数中使用
        this.handleClick = this.handleClick.bind(this);
    }

    componentDidMount() {
        console.log("execute right now")
    }

    handleClick(name, event) {
        console.log(name)
        this.setState(prevState => ({
            date: new Date()
        }));
        console.log(this.state)
        utilsBox.httpRequest("get", "/files/xxx?id=123654", {"aaa": "bbb"})
    }

    render() {
        return (
            <ul className="nav">
                <li className="nav-item">
                    <a className="nav-link active" href="#" onClick={this.handleClick.bind(this, "Name-A")}>目录A</a>
                </li>
                <li className="nav-item">
                    <a className="nav-link" href="#" onClick={this.handleClick.bind(this, "Name-B")}>目录B</a>
                </li>
                <li className="nav-item">
                    <a className="nav-link disabled">目录C {this.state.date.toLocaleTimeString()}</a>
                </li>
            </ul>
        );
    }
}

ReactDOM.render(
    <Breadcrumb/>,
    document.getElementById('breadcrumb')
);