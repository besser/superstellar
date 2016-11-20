package events

// GENERATED CODE! DO NOT EDIT THIS FILE!
// ADD YOUR EVENT AND RUN 'go generate' INSTEAD

import (
	"time"
)

const (
	buffersLength                         = 10000
	idleDispatcherSleepTime time.Duration = 5 * time.Millisecond
)

type TimeTickListener interface {
	HandleTimeTick(*TimeTick)
}

type ProjectileFiredListener interface {
	HandleProjectileFired(*ProjectileFired)
}

type ProjectileHitListener interface {
	HandleProjectileHit(*ProjectileHit)
}

type UserInputListener interface {
	HandleUserInput(*UserInput)
}

type UserJoinedListener interface {
	HandleUserJoined(*UserJoined)
}

type UserLeftListener interface {
	HandleUserLeft(*UserLeft)
}

type UserDiedListener interface {
	HandleUserDied(*UserDied)
}

type EventDispatcher struct {

	// TimeTick
	timeTickQueue     chan *TimeTick
	timeTickListeners []TimeTickListener

	// ProjectileFired
	projectileFiredQueue     chan *ProjectileFired
	projectileFiredListeners []ProjectileFiredListener

	// ProjectileHit
	projectileHitQueue     chan *ProjectileHit
	projectileHitListeners []ProjectileHitListener

	// UserInput
	userInputQueue     chan *UserInput
	userInputListeners []UserInputListener

	// UserJoined
	userJoinedQueue     chan *UserJoined
	userJoinedListeners []UserJoinedListener

	// UserLeft
	userLeftQueue     chan *UserLeft
	userLeftListeners []UserLeftListener

	// UserDied
	userDiedQueue     chan *UserDied
	userDiedListeners []UserDiedListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{

		// TimeTick
		timeTickQueue:     make(chan *TimeTick, buffersLength),
		timeTickListeners: []TimeTickListener{},

		// ProjectileFired
		projectileFiredQueue:     make(chan *ProjectileFired, buffersLength),
		projectileFiredListeners: []ProjectileFiredListener{},

		// ProjectileHit
		projectileHitQueue:     make(chan *ProjectileHit, buffersLength),
		projectileHitListeners: []ProjectileHitListener{},

		// UserInput
		userInputQueue:     make(chan *UserInput, buffersLength),
		userInputListeners: []UserInputListener{},

		// UserJoined
		userJoinedQueue:     make(chan *UserJoined, buffersLength),
		userJoinedListeners: []UserJoinedListener{},

		// UserLeft
		userLeftQueue:     make(chan *UserLeft, buffersLength),
		userLeftListeners: []UserLeftListener{},

		// UserDied
		userDiedQueue:     make(chan *UserDied, buffersLength),
		userDiedListeners: []UserDiedListener{},
	}
}

func (d *EventDispatcher) RunEventLoop() {
	for {
		select {

		// TimeTick
		case event := <-d.timeTickQueue:
			for _, listener := range d.timeTickListeners {
				listener.HandleTimeTick(event)
			}

		// ProjectileFired
		case event := <-d.projectileFiredQueue:
			for _, listener := range d.projectileFiredListeners {
				listener.HandleProjectileFired(event)
			}

		// ProjectileHit
		case event := <-d.projectileHitQueue:
			for _, listener := range d.projectileHitListeners {
				listener.HandleProjectileHit(event)
			}

		// UserInput
		case event := <-d.userInputQueue:
			for _, listener := range d.userInputListeners {
				listener.HandleUserInput(event)
			}

		// UserJoined
		case event := <-d.userJoinedQueue:
			for _, listener := range d.userJoinedListeners {
				listener.HandleUserJoined(event)
			}

		// UserLeft
		case event := <-d.userLeftQueue:
			for _, listener := range d.userLeftListeners {
				listener.HandleUserLeft(event)
			}

		// UserDied
		case event := <-d.userDiedQueue:
			for _, listener := range d.userDiedListeners {
				listener.HandleUserDied(event)
			}

		default:
			time.Sleep(idleDispatcherSleepTime)
		}
	}
}

// EVENT METHODS

// TimeTick

func (d *EventDispatcher) RegisterTimeTickListener(listener TimeTickListener) {
	d.timeTickListeners = append(d.timeTickListeners, listener)
}

func (d *EventDispatcher) FireTimeTick(e *TimeTick) {
	d.timeTickQueue <- e
}

// ProjectileFired

func (d *EventDispatcher) RegisterProjectileFiredListener(listener ProjectileFiredListener) {
	d.projectileFiredListeners = append(d.projectileFiredListeners, listener)
}

func (d *EventDispatcher) FireProjectileFired(e *ProjectileFired) {
	d.projectileFiredQueue <- e
}

// ProjectileHit

func (d *EventDispatcher) RegisterProjectileHitListener(listener ProjectileHitListener) {
	d.projectileHitListeners = append(d.projectileHitListeners, listener)
}

func (d *EventDispatcher) FireProjectileHit(e *ProjectileHit) {
	d.projectileHitQueue <- e
}

// UserInput

func (d *EventDispatcher) RegisterUserInputListener(listener UserInputListener) {
	d.userInputListeners = append(d.userInputListeners, listener)
}

func (d *EventDispatcher) FireUserInput(e *UserInput) {
	d.userInputQueue <- e
}

// UserJoined

func (d *EventDispatcher) RegisterUserJoinedListener(listener UserJoinedListener) {
	d.userJoinedListeners = append(d.userJoinedListeners, listener)
}

func (d *EventDispatcher) FireUserJoined(e *UserJoined) {
	d.userJoinedQueue <- e
}

// UserLeft

func (d *EventDispatcher) RegisterUserLeftListener(listener UserLeftListener) {
	d.userLeftListeners = append(d.userLeftListeners, listener)
}

func (d *EventDispatcher) FireUserLeft(e *UserLeft) {
	d.userLeftQueue <- e
}

// UserDied

func (d *EventDispatcher) RegisterUserDiedListener(listener UserDiedListener) {
	d.userDiedListeners = append(d.userDiedListeners, listener)
}

func (d *EventDispatcher) FireUserDied(e *UserDied) {
	d.userDiedQueue <- e
}
