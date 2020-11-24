package f1packet

/**
Motion Packet
The Motion packet gives physics Data for all the cars being driven.
There is additional Data for the car being driven with the goal of being able
to drive a Motion platform setup.

N.B. For the normalised vectors below, to convert to float values divide
by 32767.0f – 16-bit signed values are used to pack the Data and on the
assumption that direction values are always between -1.0f and 1.0f.

Frequency: Rate as specified in menus
Size: 1464 bytes (Packet size updated in Beta 3)
Version: 1
*/

const MotionPacketSize int = 1464

type carMotionData struct {
	WorldPositionX     float // World space X Position
	WorldPositionY     float // World space Y Position
	WorldPositionZ     float // World space Z Position
	WorldVelocityX     float // Velocity in world space X
	WorldVelocityY     float // Velocity in world space Y
	WorldVelocityZ     float // Velocity in world space Z
	WorldForwardDirX   int16 // World space forward X direction (normalised)
	WorldForwardDirY   int16 // World space forward Y direction (normalised)
	WorldForwardDirZ   int16 // World space forward Z direction (normalised)
	WorldRightDirX     int16 // World space right X direction (normalised)
	WorldRightDirY     int16 // World space right Y direction (normalised)
	WorldRightDirZ     int16 // World space right Z direction (normalised)
	GforceLateral      float // Lateral G-Force component
	GforceLongitudinal float // Longitudinal G-Force component
	GforceVertical     float // Vertical G-Force component
	Yaw                float // Yaw angle in radians
	Pitch              float // Pitch angle in radians
	Roll               float // Roll angle in radians
}

type Motion struct {
	CarDetails [MaxDrivers]carMotionData // Data for all cars on track

	// Extra player car ONLY Data
	SuspensionPosition     [4]float // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float // RL, RR, FL, FR
	SuspensionAcceleration [4]float // RL, RR, FL, FR
	WheelSpeed             [4]float // Speed of each wheel
	WheelSlip              [4]float // Slip ratio for each wheel
	LocalVelocityX         float    // Velocity in local space
	LocalVelocityY         float    // Velocity in local space
	LocalVelocityZ         float    // Velocity in local space
	AngularVelocityX       float    // Angular Velocity x-component
	AngularVelocityY       float    // Angular Velocity y-component
	AngularVelocityZ       float    // Angular Velocity z-component
	AngularAccelerationX   float    // Angular Velocity x-component
	AngularAccelerationY   float    // Angular Velocity y-component
	AngularAccelerationZ   float    // Angular Velocity z-component
	FrontWheelsAngle       float    // Current front wheels angle in radians
}
