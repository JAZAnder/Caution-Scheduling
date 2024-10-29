function userHour(userHour){
    this.id = userHour['id']
    this.hourId = userHour['hourId']
    this.tutor = userHour['tutor']
    this.available = userHour['available']
}

function tutorsAndHours(tutorsAndHours){
    this.id = tutorsAndHours['id']
    this.hourId = tutorsAndHours['hourId']
    this.tutorId = tutorsAndHours['tutor']
    this.firstName = tutorsAndHours['firstName']
    this.lastName = tutorsAndHours['lastName']
    this.startTime = tutorsAndHours['startTime']
    this.endTime = tutorsAndHours['endTime']
    this.dayOfWeek = tutorsAndHours['dayOfWeek']
}

export default  userHour