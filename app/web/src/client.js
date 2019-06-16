export const VALVE_WARM = 'warm';
export const VALVE_OSMOSE = 'osmose';
export const VALVE_COLD = 'cold'; // window.setInterval(() => {, refreshInterval

export class WasserSpenderClient {
  constructor(basePath) {
    this.basePath = basePath || "";
  }

  findTimer() {
    return fetch(this.basePath + '/api/v1/timer/').then(res => res.json()).then(data => {
      let time = new Date(new Date(data.End) - new Date());
      return time;
    });
  }

  findAllValves() {
    return fetch(this.basePath + `/api/v1/valves/`).then(res => res.json());
  }

  setValve(name, open) {
    return fetch(this.basePath + `/api/v1/valves/${name}`, {
      method: 'PUT',
      body: JSON.stringify({
        Open: open
      })
    });
  }

}